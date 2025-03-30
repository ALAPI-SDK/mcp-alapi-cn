package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"mcp-alapi-cn/internal/models"

	"github.com/go-resty/resty/v2"
	"github.com/mark3labs/mcp-go/mcp"
)

type ToolHandler struct {
	baseURL string
	token   string
	client  *resty.Client
}

func NewToolHandler(baseURL string, token string) *ToolHandler {
	return &ToolHandler{
		baseURL: baseURL,
		token:   token,
		client:  resty.New(),
	}
}

func (h *ToolHandler) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.Params.Arguments
	uri := request.Params.Name

	marshal, err := json.Marshal(arguments)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}

	response, err := h.client.SetBaseURL(h.baseURL).R().
		SetHeader("Content-Type", "application/json").
		SetHeader("token", h.token).
		SetHeader("User-Agent", "ALAPI-SDK/MCP-SERVER v1.0.0").
		SetBody(marshal).
		Post(uri)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}

	// 解析基础响应
	var baseResp models.BaseResponse
	if err := json.Unmarshal(response.Body(), &baseResp); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}
	if baseResp.Code != 200 {
		return nil, fmt.Errorf("api response failed:%s", baseResp.Message)
	}

	// 只返回 data 部分的数据
	dataJSON, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, fmt.Errorf("marshal response data failed: %w", err)
	}

	return mcp.NewToolResultText(string(dataJSON)), nil
}
