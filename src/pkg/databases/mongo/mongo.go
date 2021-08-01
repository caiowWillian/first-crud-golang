package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Insert(operation MongoOperation) error
}

type MongoOperation struct {
	Database   string
	Collection string
	Value      interface{}
}

type repository struct {
	db  *mongo.Client
	err error
}

var repo Repository

func newMongo() Repository {
	options := options.Client().ApplyURI("mongodb://root:root@localhost:8081/")
	client, err := mongo.Connect(context.TODO(), options)

	return &repository{client, err}
}

func init() {

	options := options.Client().ApplyURI("mongodb://root:root@localhost:8081/")
	client, err := mongo.Connect(context.TODO(), options)

	repo = &repository{client, err}
}

func (repo *repository) Insert(operation MongoOperation) error {
	_, err := repo.db.Database(operation.Database).Collection(operation.Collection).InsertOne(context.TODO(), operation.Value)
	return err
}

func Repo() Repository {
	return repo
}
