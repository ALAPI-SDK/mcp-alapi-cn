package openapi

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-resty/resty/v2"
)

type Loader struct {
	ctx   context.Context
	token string
}

func NewLoader(ctx context.Context, token string) *Loader {
	return &Loader{ctx: ctx, token: token}
}

func (l *Loader) LoadSpec(specURL string) (*openapi3.T, error) {
	loader := &openapi3.Loader{Context: l.ctx, IsExternalRefsAllowed: true}

	client := resty.New()

	// 传递 TOKEN 请求头, 用于获取有权限的接口列表
	response, err := client.R().SetHeader("token", l.token).Get(specURL)
	if err != nil {
		return nil, err
	}

	return loader.LoadFromData(response.Body())
}
