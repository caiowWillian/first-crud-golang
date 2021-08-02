package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Insert(operation MongoOperation) error
	Ping() error
	GetAll(operation MongoOperation, data interface{}) error
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
	options := options.Client().ApplyURI("mongodb://localhost:27017/teste")
	client, err := mongo.Connect(context.TODO(), options)

	return &repository{client, err}
}

func init() {

	options := options.Client().ApplyURI("mongodb://localhost:27017/teste")
	client, err := mongo.Connect(context.TODO(), options)

	repo = &repository{client, err}
}

func (repo *repository) Insert(operation MongoOperation) error {
	_, err := repo.db.Database(operation.Database).Collection(operation.Collection).InsertOne(context.TODO(), operation.Value)
	return err
}

func (repo *repository) Ping() error {
	err := repo.db.Ping(context.TODO(), nil)

	return err
}

func (repo *repository) GetAll(operation MongoOperation, data interface{}) error {
	cursor, err := repo.db.Database(operation.Database).Collection(operation.Collection).Find(context.Background(), bson.D{})

	if err != nil {
		return err
	}
	cursor.All(context.Background(), data)

	if err != nil {
		return err
	}
	return nil
}

func Repo() Repository {
	return repo
}
