package mongo

import (
    "context"

    "go.mongodb.org/mongo-driver/mongo"
    "{{.NameSpace}}/internal/domain/entity"
    "{{.NameSpace}}/internal/domain/repository"
)

const (
    collectionName = "{{.Domain}}"
)

type {{.Domain}}RepositoryImpl struct {
    db *mongo.Database
}

func (r *{{.Domain}}RepositoryImpl) Save(ctx context.Context, data *entity.{{.Domain}}) error {
    panic("TODO: Implement")
}

func (r *{{.Domain}}RepositoryImpl) Update(ctx context.Context, data *entity.{{.Domain}}) error {
    panic("TODO: Implement")
}

func (r *{{.Domain}}RepositoryImpl) Delete(ctx context.Context, id string) error {
    panic("TODO: Implement")
}

func (r *{{.Domain}}RepositoryImpl) Find(ctx context.Context) ([]entity.{{.Domain}}, error) {
    panic("TODO: Implement")
}

func (r *{{.Domain}}RepositoryImpl) FindOne(ctx context.Context) (*entity.{{.Domain}}, error) {
    panic("TODO: Implement")
}

func New{{.Domain}}Repository(db *mongo.Database) repository.{{.Domain}}Repository {
    return &{{.Domain}}RepositoryImpl{
        db: db,
    }
}