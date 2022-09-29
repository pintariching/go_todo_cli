package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Description  string             `bson:"description,omitempty"`
	Done         bool               `bson:"done,omitempty"`
	Created_at   time.Time          `bson:"created_at,omitempty"`
	Completed_at time.Time          `bson:"updated_at,omitempty"`
}

func InsertTodo(name string, col *mongo.Collection, ctx context.Context) error {
	todo := Todo{
		Name:       name,
		Done:       false,
		Created_at: time.Now(),
	}

	_, err := col.InsertOne(ctx, todo)

	if err != nil {
		return err
	}

	return nil
}

func UpdateWithDescription(name string, description string, col *mongo.Collection, ctx context.Context) {
	todo := bson.M{
		"$set": bson.M{"description": description},
	}

	col.FindOneAndUpdate(ctx, bson.M{"name": name}, todo)
}

func Complete(name string, col *mongo.Collection, ctx context.Context) {
	todo := bson.M{
		"$set": bson.M{"done": true},
	}

	col.FindOneAndUpdate(ctx, bson.M{"name": name}, todo)
}

func NotComplete(name string, col *mongo.Collection, ctx context.Context) {
	todo := bson.M{
		"$set": bson.M{"done": false},
	}

	col.FindOneAndUpdate(ctx, bson.M{"name": name}, todo)
}

func Delete(name string, col *mongo.Collection, ctx context.Context) error {
	_, err := col.DeleteOne(ctx, bson.M{"name": name})

	if err != nil {
		return err
	}

	return nil
}
