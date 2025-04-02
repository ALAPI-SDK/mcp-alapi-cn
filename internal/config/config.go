package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	ServerName string
	Version    string
	OpenAPIURL string
	BaseURL    string
	Token      string
}

// getVersion 从 Git tag 获取版本号
func getVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(strings.TrimPrefix(string(output), "v"))
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
		Version:    getVersion(),
		OpenAPIURL: openAPIURL,
		BaseURL:    "https://v3.alapi.cn",
		Token:      token,
	}, nil
}
