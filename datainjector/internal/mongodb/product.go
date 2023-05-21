package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Product struct {
	EventId    primitive.ObjectID `bson:"_id,omitempty"`
	ProductId  int                `bson:"product_id"`
	Name       string             `bson:"name,omitempty"`
	Price      int                `bson:"price,omitempty"`
	ModifiedAt time.Time          `bson:"modified_at"`
}

type ProductEventRepository struct {
	Collection *mongo.Collection
}

func NewProductStreamRepository(m *MongoDB) *ProductEventRepository {
	collection := m.Database.Collection("product_events")
	return &ProductEventRepository{
		Collection: collection,
	}
}

func (r *ProductEventRepository) Insert(product *Product) error {
	result, err := r.Collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	id := result.InsertedID.(primitive.ObjectID)
	product.EventId = id
	return nil
}

func (r *ProductEventRepository) Count() (int64, error) {
	return r.Collection.CountDocuments(context.TODO(), bson.D{})
}
