package main

import (
	"context"
	"fmt"
	"log"
	"mcp-alapi-cn/internal/config"
	"mcp-alapi-cn/internal/server"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	srv := server.NewServer(cfg)
	if err := srv.Initialize(ctx); err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	fmt.Printf("MCP Server initialized with OpenAPI spec from: %s\n", cfg.OpenAPIURL)
	if err := srv.Serve(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
