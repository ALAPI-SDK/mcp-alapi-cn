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

	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 将配置添加到 context 中
	ctx = config.WithConfig(ctx, cfg)

	srv := server.NewServer(ctx, cfg)

	// 初始化OpenAPI工具
	if err := srv.InitializeOpenAPI(ctx); err != nil {
		log.Fatalf("初始化OpenAPI工具失败: %v", err)
	}

	// 初始化自定义工具
	srv.InitializeCustomTool(ctx)

	fmt.Printf("MCP Server initialized with OpenAPI spec from: %s\n", cfg.OpenAPIURL)
	if err := srv.Start(); err != nil {
		log.Fatalf("服务器退出: %v", err)
	}
}
