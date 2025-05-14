package port

import (
	"main/internal/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetById(id string) (model.Todo, error)
	Create(entry model.Todo) error
	Update(id string, entry model.Todo) (*mongo.UpdateResult, error)
	Delete(id string) error
}
