package server

import (
	"context"
	"flag"
	"fmt"
	"mcp-alapi-cn/internal/config"
	"mcp-alapi-cn/internal/handler"
	"mcp-alapi-cn/internal/openapi"
	"mcp-alapi-cn/internal/tools"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	mcpServer *server.MCPServer
	config    *config.Config
	handler   *handler.OpenAPIToolHandler
	ctx       context.Context
}

func NewServer(ctx context.Context, cfg *config.Config) *Server {

	return &Server{
		mcpServer: server.NewMCPServer(cfg.ServerName, cfg.Version),
		config:    cfg,
		handler:   handler.NewOpenAPIToolHandler(cfg.BaseURL, cfg.Token),
		ctx:       ctx,
	}
}

func (s *Server) InitializeOpenAPI(ctx context.Context) error {

	loader := openapi.NewLoader(ctx, s.config.Token)
	doc, err := loader.LoadSpec(s.config.OpenAPIURL)
	if err != nil {
		return fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	return s.registerOpenAPITools(doc)
}

func (s *Server) InitializeCustomTool(ctx context.Context) {

	tools.RegisterTools(ctx, s.mcpServer)
}

func (s *Server) wrapHandler(handler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		ctxWithConfig := config.WithConfig(ctx, s.config)
		return handler(ctxWithConfig, req)
	}
}

func (s *Server) registerOpenAPITools(doc *openapi3.T) error {
	toolCount := 0
	for path, item := range doc.Paths.Map() {
		if item.Post != nil {
			tool := mcp.NewTool(path, mcp.WithDescription(item.Post.Summary))
			schema := item.Post.RequestBody.Value.Content["application/json"].Schema

			requiredParams := make(map[string]bool)
			for _, required := range schema.Value.Required {
				requiredParams[required] = true
			}

			for paramName, ref := range schema.Value.Properties {
				description := ref.Value.Description
				if requiredParams[paramName] {
					mcp.WithString(paramName, mcp.Description(description), mcp.Required())(&tool)
				} else {
					mcp.WithString(paramName, mcp.Description(description))(&tool)
				}
			}

			s.mcpServer.AddTool(tool, s.wrapHandler(s.handler.Handle))
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

func (s *Server) Start() error {
	return server.ServeStdio(s.mcpServer)
}
