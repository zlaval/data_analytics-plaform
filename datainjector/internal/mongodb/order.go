package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrderDetail struct {
	Data string
}

type Order struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	ProductId int                `bson:"product_id"`
	UserId    int                `bson:"user_id"`
	TimeStamp time.Time          `bson:"timestamp"`
	Details   OrderDetail        `bson:"details"`
}

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(m *MongoDB) *OrderRepository {
	ordersCollection := m.Database.Collection("orders")

	return &OrderRepository{
		Collection: ordersCollection,
	}
}

func (r *OrderRepository) Insert(order *Order) error {
	result, err := r.Collection.InsertOne(context.TODO(), order)
	if err != nil {
		return err
	}
	id := result.InsertedID.(primitive.ObjectID)
	order.Id = id
	return nil
}

func (r *OrderRepository) Count() (int64, error) {
	return r.Collection.CountDocuments(context.TODO(), bson.D{})
}
