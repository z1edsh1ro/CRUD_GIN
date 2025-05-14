package repository

import (
	"context"
	"log"
	"main/internal/domain/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBTodoRepository struct {
	Collection *mongo.Collection
}

func NewTodoRepository(client *mongo.Client) *DBTodoRepository {
	conllection := client.Database("todo_db").Collection("todos")

	return &DBTodoRepository{
		Collection: conllection,
	}
}

func (repo *DBTodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	cursor, err := repo.Collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Panic(err)
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

func (repo *DBTodoRepository) GetById(id string) (model.Todo, error) {
	var todo model.Todo

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
		return model.Todo{}, err
	}

	err = repo.Collection.FindOne(context.Background(), bson.M{"_id": mongoId}).Decode(&todo)

	if err != nil {
		log.Panic(err)
		return model.Todo{}, err
	}

	return todo, nil
}

func (repo *DBTodoRepository) Create(entry model.Todo) error {
	_, err := repo.Collection.InsertOne(context.TODO(), model.Todo{
		Title:       entry.Title,
		Description: entry.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

func (repo *DBTodoRepository) Update(id string, entry model.Todo) (*mongo.UpdateResult, error) {
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$set", bson.D{
			{"title", entry.Title},
			{"description", entry.Description},
			{"updated_at", time.Now()},
		}},
	}

	res, err := repo.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": mongoId},
		update,
	)

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return res, nil
}

func (repo *DBTodoRepository) Delete(id string) error {
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
		return err
	}

	_, err = repo.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": mongoId},
	)

	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
