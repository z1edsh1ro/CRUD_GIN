package port

import (
	"main/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepositoryInterface interface {
	GetAll() ([]model.Todo, error)
	GetById(id primitive.ObjectID) (model.Todo, error)
	Create(entry model.Todo) error
	Update(id primitive.ObjectID, entry model.Todo) error
	Delete(id primitive.ObjectID) error
}
