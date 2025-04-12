package tools

import (
	"context"
	"mcp-alapi-cn/internal/config"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Tool 定义工具接口
type Tool interface {
	Register(ctx context.Context, s *server.MCPServer)
	Name() string
	Description() string
	// 新增设置配置的方法
	SetConfig(*config.Config)
}

// BaseTool 提供了基础工具实现
type BaseTool struct {
	Config *config.Config
}

// SetConfig 设置工具配置
func (bt *BaseTool) SetConfig(cfg *config.Config) {
	bt.Config = cfg
}

// RegisterTools 注册所有自定义工具
func RegisterTools(ctx context.Context, s *server.MCPServer) {
	// 从上下文中获取配置
	cfg, ok := config.FromContext(ctx)
	if !ok {
		panic("无法从上下文中获取配置")
	}

	tools := []Tool{
		NewUserApis(),
	}

	// 为所有工具设置配置并注册
	for _, tool := range tools {
		tool.SetConfig(cfg)
		tool.Register(ctx, s)
	}
}

// wrapToolHandler 现在简化了，仅用于确保上下文包含配置
// 主要用于外部工具的集成，内部工具通过依赖注入获取配置
func wrapToolHandler(ctx context.Context, handler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 从注册时的上下文中获取配置
	cfg, ok := config.FromContext(ctx)
	if !ok {
		// 如果没有配置，返回原始处理函数
		return handler
	}

	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// 将配置添加到调用上下文
		callCtxWithConfig := config.WithConfig(callCtx, cfg)
		// 使用带有配置的上下文调用原始处理函数
		return handler(callCtxWithConfig, req)
	}
}
