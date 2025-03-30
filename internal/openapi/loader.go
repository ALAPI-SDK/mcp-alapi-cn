package openapi

import (
	"context"
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"
)

type Loader struct {
	ctx context.Context
}

func NewLoader(ctx context.Context) *Loader {
	return &Loader{ctx: ctx}
}

func (l *Loader) LoadSpec(specURL string) (*openapi3.T, error) {
	loader := &openapi3.Loader{Context: l.ctx, IsExternalRefsAllowed: true}

	parsedURL, err := url.Parse(specURL)
	if err != nil {
		return nil, err
	}

	return loader.LoadFromURI(parsedURL)
}
