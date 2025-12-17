package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	parser := NewYApiParser("", "")

	tests := []struct {
		name           string
		url            string
		expectedBase   string
		expectedPID    string
		expectedIID    string
		expectedType   string
	}{
		{
			name:         "完整接口URL",
			url:          "http://yapi.example.com/project/123/interface/api/456",
			expectedBase: "http://yapi.example.com",
			expectedPID:  "123",
			expectedIID:  "456",
			expectedType: "interface",
		},
		{
			name:         "项目URL",
			url:          "http://yapi.example.com/project/789",
			expectedBase: "http://yapi.example.com",
			expectedPID:  "789",
			expectedIID:  "",
			expectedType: "",
		},
		{
			name:         "带查询参数的URL",
			url:          "http://yapi.example.com/interface/api/456?pid=123",
			expectedBase: "http://yapi.example.com",
			expectedPID:  "123",
			expectedIID:  "456",
			expectedType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := parser.ParseURL(tt.url)

			if info.BaseURL != tt.expectedBase {
				t.Errorf("BaseURL = %s, 期望 %s", info.BaseURL, tt.expectedBase)
			}

			if info.ProjectID != tt.expectedPID {
				t.Errorf("ProjectID = %s, 期望 %s", info.ProjectID, tt.expectedPID)
			}

			if info.InterfaceID != tt.expectedIID {
				t.Errorf("InterfaceID = %s, 期望 %s", info.InterfaceID, tt.expectedIID)
			}

			if info.Type != tt.expectedType {
				t.Errorf("Type = %s, 期望 %s", info.Type, tt.expectedType)
			}
		})
	}
}

func TestNewYApiParser(t *testing.T) {
	baseURL := "http://test.com"
	token := "test-token"

	parser := NewYApiParser(baseURL, token)

	if parser.baseURL != baseURL {
		t.Errorf("baseURL = %s, 期望 %s", parser.baseURL, baseURL)
	}

	if parser.token != token {
		t.Errorf("token = %s, 期望 %s", parser.token, token)
	}

	if parser.client == nil {
		t.Error("HTTP客户端未初始化")
	}
}

