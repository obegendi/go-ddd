package customer

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(customer *Customer) error
	Update(id string, customer *Customer) error
	Delete(customer *Customer) error
	Get(id string) (*Customer, error)
}

type repository struct {
	collection *mongo.Collection
}

func NewCustomerRepository(collection *mongo.Collection) Repository {
	return &repository{collection: collection}
}

func (r *repository) Create(customer *Customer) error {
	_, err := r.collection.InsertOne(context.Background(), customer)
	return err
}

func (r *repository) Update(id string, customer *Customer) error {
	bsonid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": bsonid}, customer)
	return err
}

func (r *repository) Delete(customer *Customer) error {
	_, err := r.collection.DeleteOne(context.Background(), customer)
	return err
}

func (r *repository) Get(id string) (*Customer, error) {
	bsonid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var customer *Customer
	err = r.collection.FindOne(context.Background(), &Customer{Id: bsonid}).Decode(&customer)
	return customer, err
}
