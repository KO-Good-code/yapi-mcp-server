package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// YApiParser YApi文档解析器
type YApiParser struct {
	baseURL string
	token   string
	client  *http.Client
}

// NewYApiParser 创建新的YApi解析器
func NewYApiParser(baseURL, token string) *YApiParser {
	return &YApiParser{
		baseURL: baseURL,
		token:   token,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// URLInfo URL解析结果
type URLInfo struct {
	BaseURL     string
	ProjectID   string
	InterfaceID string
	Type        string
}

// InterfaceInfo 接口信息结构
type InterfaceInfo struct {
	ID            int                      `json:"_id"`
	Title         string                   `json:"title"`
	Path          string                   `json:"path"`
	Method        string                   `json:"method"`
	Description   string                   `json:"desc"`
	ReqQuery      []map[string]interface{} `json:"req_query"`
	ReqBodyOther  string                   `json:"req_body_other"`
	ResBody       string                   `json:"res_body"`
	ReqHeaders    []map[string]interface{} `json:"req_headers"`
	Tag           []string                 `json:"tag"`
	Status        string                   `json:"status"`
	ProjectID     int                      `json:"project_id"`
	CatID         int                      `json:"catid"`
	UID           int                      `json:"uid"`
	AddTime       int64                    `json:"add_time"`
	UpTime        int64                    `json:"up_time"`
}

// ParseURL 解析YApi文档URL
func (p *YApiParser) ParseURL(rawURL string) URLInfo {
	info := URLInfo{}
	
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return info
	}

	info.BaseURL = fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	
	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	
	// 查找project和interface的位置
	for i, part := range pathParts {
		if part == "project" && i+1 < len(pathParts) {
			info.ProjectID = pathParts[i+1]
		}
		if part == "interface" && i+1 < len(pathParts) {
			// 支持 /interface/api/123 格式
			if pathParts[i+1] == "api" && i+2 < len(pathParts) {
				info.InterfaceID = pathParts[i+2]
				info.Type = "interface"
			}
		}
		// 支持直接的接口ID格式 /api/123
		if part == "api" && i+1 < len(pathParts) {
			info.InterfaceID = pathParts[i+1]
			info.Type = "interface"
		}
	}

	// 从查询参数获取
	query := parsedURL.Query()
	if id := query.Get("id"); id != "" {
		info.InterfaceID = id
	}
	if pid := query.Get("pid"); pid != "" {
		info.ProjectID = pid
	}

	return info
}

// GetInterfaceInfo 从YApi文档链接获取接口信息
func (p *YApiParser) GetInterfaceInfo(rawURL string) (InterfaceInfo, error) {
	urlInfo := p.ParseURL(rawURL)

	if urlInfo.InterfaceID != "" {
		return p.getInterfaceByID(urlInfo)
	}

	// 尝试从HTML解析
	return p.parseHTMLPage(rawURL)
}

// getInterfaceByID 通过接口ID获取接口信息
func (p *YApiParser) getInterfaceByID(urlInfo URLInfo) (InterfaceInfo, error) {
	// 优先使用API
	if urlInfo.ProjectID != "" {
		// 构建API URL，token作为query参数
		apiURL := fmt.Sprintf("%s/api/interface/get?id=%s", urlInfo.BaseURL, urlInfo.InterfaceID)
		if p.token != "" {
			apiURL = fmt.Sprintf("%s&token=%s", apiURL, p.token)
		}

		resp, err := p.client.Get(apiURL)
		if err != nil {
			// API请求失败，降级到HTML解析
			htmlURL := fmt.Sprintf("%s/project/%s/interface/api/%s", urlInfo.BaseURL, urlInfo.ProjectID, urlInfo.InterfaceID)
			return p.parseHTMLPage(htmlURL)
		}
			defer resp.Body.Close()
		
		if resp.StatusCode == 200 {
			var result struct {
				Errcode int           `json:"errcode"`
				Data    InterfaceInfo `json:"data"`
				Errmsg  string        `json:"errmsg"`
			}
			if json.NewDecoder(resp.Body).Decode(&result) == nil {
				if result.Errcode == 0 {
				return result.Data, nil
				}
				// API返回错误，降级到HTML解析
			}
		}
	}

	// 降级到HTML解析
	if urlInfo.ProjectID != "" && urlInfo.InterfaceID != "" {
	htmlURL := fmt.Sprintf("%s/project/%s/interface/api/%s", urlInfo.BaseURL, urlInfo.ProjectID, urlInfo.InterfaceID)
	return p.parseHTMLPage(htmlURL)
	}
	
	return InterfaceInfo{}, fmt.Errorf("无法获取接口信息：缺少必要参数或无法访问YApi")
}

// ProjectInterfaces 项目接口列表
type ProjectInterfaces struct {
	ProjectID string          `json:"project_id"`
	Interfaces []InterfaceInfo `json:"interfaces"`
	Total     int             `json:"total"`
}

// getProjectInterfaces 获取项目所有接口
func (p *YApiParser) getProjectInterfaces(urlInfo URLInfo) (ProjectInterfaces, error) {
	result := ProjectInterfaces{
		ProjectID: urlInfo.ProjectID,
		Interfaces: []InterfaceInfo{},
		Total:     0,
	}

	// 构建API URL，token作为query参数
	apiURL := fmt.Sprintf("%s/api/interface/list?project_id=%s&page=1&limit=1000", urlInfo.BaseURL, urlInfo.ProjectID)
	if p.token != "" {
		apiURL = fmt.Sprintf("%s&token=%s", apiURL, p.token)
	}

	resp, err := p.client.Get(apiURL)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			var apiResult struct {
				Errcode int `json:"errcode"`
				Data    struct {
					List []InterfaceInfo `json:"list"`
				} `json:"data"`
			Errmsg string `json:"errmsg"`
			}
		if json.NewDecoder(resp.Body).Decode(&apiResult) == nil {
			if apiResult.Errcode == 0 {
				result.Interfaces = apiResult.Data.List
				result.Total = len(result.Interfaces)
				return result, nil
			}
			return result, fmt.Errorf("YApi API错误: %s (errcode: %d)", apiResult.Errmsg, apiResult.Errcode)
		}
	}

	if err != nil {
		return result, fmt.Errorf("无法连接到YApi: %v", err)
	}
	return result, fmt.Errorf("无法获取项目接口，请检查token和项目ID")
}

// parseHTMLPage 从HTML页面解析接口信息
func (p *YApiParser) parseHTMLPage(pageURL string) (InterfaceInfo, error) {
	resp, err := p.client.Get(pageURL)
	if err != nil {
		return InterfaceInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return InterfaceInfo{}, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return InterfaceInfo{}, err
	}

	// 尝试从script标签中提取JSON数据
	htmlContent := string(body)
	re := regexp.MustCompile(`(?s)interfaceData\s*[:=]\s*(\{.*?\})`)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) > 1 {
		var data InterfaceInfo
		if json.Unmarshal([]byte(matches[1]), &data) == nil {
			return data, nil
		}
	}

	// 从HTML元素提取
	return p.extractFromHTML(htmlContent)
}

// extractFromHTML 从HTML元素提取接口信息
func (p *YApiParser) extractFromHTML(htmlContent string) (InterfaceInfo, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return InterfaceInfo{}, err
	}

	info := InterfaceInfo{
		ReqQuery:     []map[string]interface{}{},
		ReqBodyOther: "",
		ResBody:      "",
		ReqHeaders:   []map[string]interface{}{},
		Tag:          []string{},
	}

	var extractText func(*html.Node, *strings.Builder)
	extractText = func(n *html.Node, buf *strings.Builder) {
		if n.Type == html.TextNode {
			buf.WriteString(strings.TrimSpace(n.Data))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractText(c, buf)
		}
	}

	// 查找标题
	var titleBuilder strings.Builder
	var findTitle func(*html.Node)
	findTitle = func(n *html.Node) {
		if n.Data == "h1" || (n.Type == html.ElementNode && hasClass(n, "title")) {
			extractText(n, &titleBuilder)
			info.Title = titleBuilder.String()
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTitle(c)
		}
	}
	findTitle(doc)

	return info, nil
}

// hasClass 检查节点是否有指定class
func hasClass(n *html.Node, className string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			classes := strings.Fields(attr.Val)
			for _, c := range classes {
				if c == className {
					return true
				}
			}
		}
	}
	return false
}

