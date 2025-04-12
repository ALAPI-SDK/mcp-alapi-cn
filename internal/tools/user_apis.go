package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-alapi-cn/internal/models"

	"github.com/go-resty/resty/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type UserApis struct {
	BaseTool // 嵌入 BaseTool 来继承配置功能
}

func NewUserApis() *UserApis {
	return &UserApis{}
}

func (c *UserApis) Name() string {
	return "user_apis"
}

func (c *UserApis) Description() string {
	return "获取已申请的API接口"
}

func (c *UserApis) Register(ctx context.Context, s *server.MCPServer) {
	tool := mcp.NewTool(c.Name(),
		mcp.WithDescription(c.Description()),
	)

	// 不需要再使用 wrapToolHandler，直接使用自己的处理函数
	s.AddTool(tool, c.Handle)
}

func (c *UserApis) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 直接使用嵌入的配置，不需要从上下文获取
	if c.Config == nil {
		return nil, fmt.Errorf("工具配置未设置")
	}

	response, err := resty.New().SetBaseURL(c.Config.BaseURL).R().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", c.Config.Token).
		SetHeader("User-Agent", "ALAPI-SDK/MCP-SERVER v1.0.0").
		Get("/api/user_apis")
	if err != nil {
		return nil, fmt.Errorf("API请求失败: %w", err)
	}

	var baseResp models.BaseResponse
	if err := json.Unmarshal(response.Body(), &baseResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	if baseResp.Code != 200 {
		return nil, fmt.Errorf("请求参数错误: %s", baseResp.Message)
	}

	dataJSON, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, fmt.Errorf("序列化响应数据失败: %w", err)
	}

	return mcp.NewToolResultText(string(dataJSON)), nil
}
