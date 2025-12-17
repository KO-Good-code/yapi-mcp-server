package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// MCP服务器实现
type MCPServer struct {
	stdin  io.Reader
	stdout io.Writer
}

// 接口信息结构（定义在yapi_parser.go中）

// MCP消息处理
func (s *MCPServer) handleRequest(request map[string]interface{}) map[string]interface{} {
	method, ok := request["method"].(string)
	if !ok {
		return map[string]interface{}{
			"result": map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": "Invalid request: missing method",
					},
				},
			},
		}
	}

	switch method {
	case "initialize":
		return map[string]interface{}{
			"result": map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"capabilities": map[string]interface{}{
					"tools": map[string]interface{}{},
				},
				"serverInfo": map[string]interface{}{
					"name":    "yapi-mcp-server",
					"version": "1.0.0",
				},
			},
		}
	case "tools/list":
		return s.handleListTools()
	case "tools/call":
		result := s.handleCallTool(request)
		// tools/call 返回的结果已经包含 content 或 error，直接返回
		return result
	default:
		return map[string]interface{}{
			"result": map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("Method not found: %s", method),
					},
				},
			},
		}
	}
}

func (s *MCPServer) handleListTools() map[string]interface{} {
	return map[string]interface{}{
		"result": map[string]interface{}{
		"tools": []map[string]interface{}{
			{
				"name":        "get_yapi_interface",
				"description": "从YApi文档链接读取接口信息",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"url": map[string]interface{}{
							"type":        "string",
							"description": "YApi文档链接",
						},
					},
					"required": []string{"url"},
				},
			},
			{
				"name":        "get_yapi_project_interfaces",
				"description": "获取YApi项目中所有接口的列表",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"project_url": map[string]interface{}{
							"type":        "string",
							"description": "YApi项目URL或项目ID",
						},
					},
					"required": []string{"project_url"},
					},
				},
			},
		},
	}
}

func (s *MCPServer) handleCallTool(request map[string]interface{}) map[string]interface{} {
	params, ok := request["params"].(map[string]interface{})
	if !ok {
		return map[string]interface{}{
			"result": map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": "Invalid params",
					},
				},
			},
		}
	}

	toolName, ok := params["name"].(string)
	if !ok {
		return map[string]interface{}{
			"result": map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": "Missing tool name",
			},
				},
			},
		}
		}

	arguments, ok := params["arguments"].(map[string]interface{})
	if !ok {
		arguments = make(map[string]interface{})
	}

	baseURL := os.Getenv("YAPI_BASE_URL")
	token := os.Getenv("YAPI_TOKEN")
	parser := NewYApiParser(baseURL, token)

	switch toolName {
	case "get_yapi_interface":
		url, ok := arguments["url"].(string)
		if !ok {
			return map[string]interface{}{
				"result": map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": "Missing url parameter",
						},
					},
				},
			}
		}
		info, err := parser.GetInterfaceInfo(url)
		if err != nil {
			return map[string]interface{}{
				"result": map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": fmt.Sprintf("Error: %s", err.Error()),
						},
					},
				},
			}
		}
		data, _ := json.MarshalIndent(info, "", "  ")
		return map[string]interface{}{
			"result": map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": string(data),
					},
				},
			},
		}

	case "get_yapi_project_interfaces":
		projectURL, ok := arguments["project_url"].(string)
		if !ok {
			return map[string]interface{}{
				"result": map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": "Missing project_url parameter",
						},
					},
				},
			}
		}
		urlInfo := parser.ParseURL(projectURL)
		if urlInfo.ProjectID == "" {
			urlInfo.ProjectID = projectURL // 假设直接是项目ID
		}
		if urlInfo.BaseURL == "" {
			urlInfo.BaseURL = baseURL
			if urlInfo.BaseURL == "" {
				urlInfo.BaseURL = "http://localhost"
			}
		}
		info, err := parser.getProjectInterfaces(urlInfo)
		if err != nil {
			return map[string]interface{}{
				"result": map[string]interface{}{
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": fmt.Sprintf("Error: %s", err.Error()),
						},
					},
				},
			}
		}
		data, _ := json.MarshalIndent(info, "", "  ")
		return map[string]interface{}{
			"result": map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": string(data),
					},
				},
			},
		}

	default:
		return map[string]interface{}{
			"result": map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("Unknown tool: %s", toolName),
					},
				},
			},
		}
	}
}

func runMCPServer() {
	server := &MCPServer{
		stdin:  os.Stdin,
		stdout: os.Stdout,
	}

	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		var request map[string]interface{}
		if err := decoder.Decode(&request); err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("解码错误: %v", err)
			continue
		}

		// 检查是否是通知（notification）- 通知没有 id 字段，不需要响应
		_, hasID := request["id"]
		if !hasID {
			// 这是通知，不需要响应，直接跳过
			method, _ := request["method"].(string)
			if method == "notifications/initialized" {
				// 初始化通知，忽略即可
				continue
			}
			// 其他通知也忽略
			continue
		}

		// 这是请求（request），需要响应
		response := server.handleRequest(request)
		
		// 确保响应包含必需的字段
		response["jsonrpc"] = "2.0"
		if id, ok := request["id"]; ok {
			response["id"] = id
		} else {
			// 理论上不应该到这里，因为上面已经检查了
			response["id"] = 0
		}

		if err := encoder.Encode(response); err != nil {
			log.Printf("编码错误: %v", err)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "yapi-mcp-server",
	Short: "YApi MCP Server - 从YApi文档链接读取接口信息",
	Long:  `一个Model Context Protocol服务器，用于从YApi文档链接读取接口信息并提供给大模型使用。`,
	Run: func(cmd *cobra.Command, args []string) {
		runMCPServer()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}


