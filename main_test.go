package main

import (
	"os"
	"testing"
)

func TestHandleListTools(t *testing.T) {
	server := &MCPServer{}
	response := server.handleListTools()

	if response["tools"] == nil {
		t.Error("tools字段不存在")
		return
	}

	tools, ok := response["tools"].([]map[string]interface{})
	if !ok {
		t.Error("tools字段类型错误")
		return
	}

	if len(tools) == 0 {
		t.Error("工具列表为空")
		return
	}

	// 检查工具名称
	found := false
	for _, tool := range tools {
		if name, ok := tool["name"].(string); ok {
			if name == "get_yapi_interface" {
				found = true
				break
			}
		}
	}

	if !found {
		t.Error("未找到get_yapi_interface工具")
	}
}

func TestHandleCallTool_GetInterface(t *testing.T) {
	// 设置测试环境变量
	os.Setenv("YAPI_BASE_URL", "http://test.example.com")
	os.Setenv("YAPI_TOKEN", "test-token")
	defer os.Unsetenv("YAPI_BASE_URL")
	defer os.Unsetenv("YAPI_TOKEN")

	server := &MCPServer{}
	request := map[string]interface{}{
		"params": map[string]interface{}{
			"name": "get_yapi_interface",
			"arguments": map[string]interface{}{
				"url": "http://test.example.com/project/123/interface/api/456",
			},
		},
	}

	response := server.handleCallTool(request)

	// 检查是否有错误或内容
	if response["error"] == nil && response["content"] == nil {
		t.Error("响应中既没有error也没有content")
	}
}

func TestHandleCallTool_MissingParam(t *testing.T) {
	server := &MCPServer{}
	request := map[string]interface{}{
		"params": map[string]interface{}{
			"name": "get_yapi_interface",
			"arguments": map[string]interface{}{},
		},
	}

	response := server.handleCallTool(request)

	if response["error"] == nil {
		t.Error("应该返回错误，因为缺少url参数")
	}
}

func TestHandleRequest_Initialize(t *testing.T) {
	server := &MCPServer{}
	request := map[string]interface{}{
		"method": "initialize",
		"params": map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":     map[string]interface{}{},
		},
	}

	response := server.handleRequest(request)

	if response["result"] == nil {
		t.Error("初始化响应应该包含result字段")
	}
}

func TestHandleRequest_UnknownMethod(t *testing.T) {
	server := &MCPServer{}
	request := map[string]interface{}{
		"method": "unknown_method",
	}

	response := server.handleRequest(request)

	if response["error"] == nil {
		t.Error("未知方法应该返回错误")
	}
}

