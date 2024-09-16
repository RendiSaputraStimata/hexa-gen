package service

import (
    "context"

    "{{.NameSpace}}/internal/application/port"
	"{{.NameSpace}}/internal/application/port/in"
	"{{.NameSpace}}/internal/application/port/out"
	"{{.NameSpace}}/internal/domain/repository"
)

type {{.Domain}}ServiceImpl struct {
    repo repository.{{.Domain}}Repository
}

func (s *{{.Domain}}ServiceImpl) Create(ctx context.Context, request *in.Create{{.Domain}}Request) (*out.{{.Domain}}CreatedResponse, error) {
    panic("NO: implement")
}

func (s *{{.Domain}}ServiceImpl) Update(ctx context.Context, request *in.Update{{.Domain}}Request) (*out.{{.Domain}}Response, error) {
    panic("NO: implement")
}

func (s *{{.Domain}}ServiceImpl) GetByID(ctx context.Context, id string) (*out.{{.Domain}}StandardResponse, error) {
    panic("NO: implement")
}

func (s *{{.Domain}}ServiceImpl) GetWithFilter(ctx context.Context, filter any) (*out.{{.Domain}}ListResponse, error) {
    panic("NO: implement")
}

func New{{.Domain}}Service(
    repo repository.{{.Domain}}Repository,
) port.{{.Domain}}Service {
    return &{{.Domain}}ServiceImpl{
        repo: repo,
    }
}