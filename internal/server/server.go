package server

import (
	"context"
	"flag"
	"fmt"
	"mcp-alapi-cn/internal/config"
	"mcp-alapi-cn/internal/handler"
	"mcp-alapi-cn/internal/openapi"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	mcpServer *server.MCPServer
	config    *config.Config
	handler   *handler.ToolHandler
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		mcpServer: server.NewMCPServer(cfg.ServerName, cfg.Version),
		config:    cfg,
		handler:   handler.NewToolHandler(cfg.BaseURL, cfg.Token),
	}
}

func (s *Server) Initialize(ctx context.Context) error {
	loader := openapi.NewLoader(ctx)
	doc, err := loader.LoadSpec(s.config.OpenAPIURL)
	if err != nil {
		return fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	return s.registerTools(doc)
}

func (s *Server) registerTools(doc *openapi3.T) error {
	toolCount := 0
	for path, item := range doc.Paths.Map() {
		if item.Post != nil {
			tool := mcp.NewTool(path, mcp.WithDescription(item.Post.Summary))
			schema := item.Post.RequestBody.Value.Content["application/json"].Schema

			// 获取必填参数列表
			requiredParams := make(map[string]bool)
			for _, required := range schema.Value.Required {
				requiredParams[required] = true
			}

			// 处理所有参数
			for paramName, ref := range schema.Value.Properties {
				description := ref.Value.Description
				if requiredParams[paramName] {
					mcp.WithString(paramName, mcp.Description(description), mcp.Required())(&tool)
				} else {
					mcp.WithString(paramName, mcp.Description(description))(&tool)
				}
			}

			s.mcpServer.AddTool(tool, s.handler.Handle)
			toolCount++
		}
	}

	if toolCount == 0 {
		return fmt.Errorf("no tools were registered from the OpenAPI spec")
	}

	return nil
}

func (s *Server) Serve() error {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio or sse)")
	flag.Parse()

	if transport == "sse" {
		sseServer := server.NewSSEServer(s.mcpServer, server.WithBaseURL(":8080"))
		if err := sseServer.Start(":8080"); err != nil {
			return err
		}
		return nil
	}

	return server.ServeStdio(s.mcpServer)
}
