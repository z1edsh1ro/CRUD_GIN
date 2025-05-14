package port

import "main/internal/domain/model"

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetById(id int) (model.Todo, error)
	Create(todo model.Todo) error
	Update(todo model.Todo) error
	Delete(id int) error
}
