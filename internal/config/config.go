package config

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 定义 context key 类型，避免与其他包的 key 冲突
type contextKey string

// ConfigKey 是在 context 中存储 Config 的 key
const ConfigKey contextKey = "config"

type Config struct {
	ServerName string `yaml:"server_name"`
	Version    string `yaml:"version"`
	BaseURL    string `yaml:"base_url"`
	Token      string `yaml:"token"`
	OpenAPIURL string `yaml:"openapi_url"`
}

// getVersion 从 Git tag 获取版本号
func getVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		return "latest"
	}
	return strings.TrimSpace(strings.TrimPrefix(string(output), "v"))
}

// WithConfig 返回带有配置信息的新 context
func WithConfig(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, ConfigKey, cfg)
}

// FromContext 从 context 中获取配置信息
func FromContext(ctx context.Context) (*Config, bool) {
	cfg, ok := ctx.Value(ConfigKey).(*Config)
	return cfg, ok
}

// MustFromContext 从 context 中获取配置信息，如果不存在则 panic
func MustFromContext(ctx context.Context) *Config {
	cfg, ok := FromContext(ctx)
	if !ok {
		panic("config not found in context")
	}
	return cfg
}

func NewConfig(ctx context.Context) (*Config, error) {
	token := os.Getenv("ALAPI_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("ALAPI_TOKEN environment variable is required")
	}

	apiID := os.Getenv("ALAPI_API_ID")
	openAPIURL := "https://v3.alapi.cn/openapi.json"
	if apiID != "" && apiID != "0" {
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
