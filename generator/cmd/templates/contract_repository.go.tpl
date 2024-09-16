package repository

import (
	"context"

	"{{.NameSpace}}/internal/domain/entity"
)

type {{.Domain}}Repository interface {
	Save(ctx context.Context, data *entity.{{.Domain}}) error
	Update(ctx context.Context, data *entity.{{.Domain}}) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context) ([]entity.{{.Domain}}, error)
	FindOne(ctx context.Context) (*entity.{{.Domain}}, error)
}