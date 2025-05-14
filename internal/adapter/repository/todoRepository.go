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
	Client *mongo.Client
}

func NewTodoRepository(mongo *mongo.Client) *DBTodoRepository {
	return &DBTodoRepository{
		Client: mongo,
	}
}

func getTodoCollection(repo *DBTodoRepository) *mongo.Collection {
	return repo.Client.Database("todo_db").Collection("todos")
}

func (repo *DBTodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	collection := getTodoCollection(repo)

	cursor, err := collection.Find(context.TODO(), bson.D{})

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

	collection := getTodoCollection(repo)

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
		return model.Todo{}, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoId}).Decode(&todo)

	if err != nil {
		log.Panic(err)
		return model.Todo{}, err
	}

	return todo, nil
}

func (repo *DBTodoRepository) Create(entry model.Todo) error {
	collection := getTodoCollection(repo)

	_, err := collection.InsertOne(context.TODO(), model.Todo{
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
	collection := getTodoCollection(repo)
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

	res, err := collection.UpdateOne(
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
	collection := getTodoCollection(repo)
	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
		return err
	}

	_, err = collection.DeleteOne(
		context.Background(),
		bson.M{"_id": mongoId},
	)

	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
