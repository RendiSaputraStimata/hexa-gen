package port

import (
	"context"

	"{{.NameSpace}}/internal/application/port/in"
	"{{.NameSpace}}/internal/application/port/out"
)

type {{.Domain}}Service interface {
	Create(ctx context.Context, request *in.Create{{.Domain}}Request) (*out.{{.Domain}}CreatedResponse, error)
	Update(ctx context.Context, request *in.Update{{.Domain}}Request) (*out.{{.Domain}}Response, error)
	GetByID(ctx context.Context, id string) (*out.{{.Domain}}StandardResponse, error)
    GetWithFilter(ctx context.Context, filter any) (*out.{{.Domain}}ListResponse, error)
}
