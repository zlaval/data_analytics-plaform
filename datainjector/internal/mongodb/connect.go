package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoConfig struct {
	Uri string
}

type MongoDB struct {
	client   *mongo.Client
	cancel   context.CancelFunc
	context  context.Context
	Database *mongo.Database
}

func (m *MongoConfig) Connect() *MongoDB {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.Uri))

	if err != nil {
		log.Panic(err)
	}

	database := client.Database("admin")

	return &MongoDB{
		client:   client,
		cancel:   cancel,
		context:  ctx,
		Database: database,
	}
}

func (m *MongoDB) Close() {
	err := m.client.Disconnect(m.context)
	if err != nil {
		log.Panic(err)
	}
}

func (m *MongoDB) Test() {
	err := m.client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println("Cannot ping the db")
	} else {
		log.Println("Ping was successful")
	}
}
