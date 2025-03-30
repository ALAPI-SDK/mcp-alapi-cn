package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerName string
	Version    string
	OpenAPIURL string
	BaseURL    string
	Token      string
}

func NewConfig() (*Config, error) {
	token := os.Getenv("ALAPI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("ALAPI_TOKEN environment variable is required")
	}

	apiID := os.Getenv("ALAPI_API_ID")
	openAPIURL := "https://v3.alapi.cn/openapi.json"
	if apiID != "" {
		openAPIURL = fmt.Sprintf("https://v3.alapi.cn/openapi/%s.json", apiID)
	}

	return &Config{
		ServerName: "ALAPI MCP Server",
		Version:    "1.0.0",
		OpenAPIURL: openAPIURL,
		BaseURL:    "https://v3.alapi.cn",
		Token:      token,
	}, nil
}
