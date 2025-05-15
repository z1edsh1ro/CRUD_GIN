package repository

import (
	"context"
	"main/internal/domain/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoCollection struct {
	Collection *mongo.Collection
}

func NewTodoRepository(client *mongo.Client) *TodoCollection {
	conllection := client.Database("todo_db").Collection("todos")

	return &TodoCollection{Collection: conllection}
}

func (repo *TodoCollection) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	cursor, err := repo.Collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *TodoCollection) GetById(id primitive.ObjectID) (model.Todo, error) {
	var todo model.Todo

	err := repo.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)

	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func (repo *TodoCollection) Create(entry model.Todo) error {
	_, err := repo.Collection.InsertOne(context.TODO(), model.Todo{
		Title:       entry.Title,
		Description: entry.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return err
}

func (repo *TodoCollection) Update(id primitive.ObjectID, entry model.Todo) error {
	update := bson.D{
		{"$set", bson.D{
			{"title", entry.Title},
			{"description", entry.Description},
			{"updated_at", time.Now()},
		}},
	}

	_, err := repo.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	return err
}

func (repo *TodoCollection) Delete(id primitive.ObjectID) error {
	_, err := repo.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": id},
	)

	return err
}
