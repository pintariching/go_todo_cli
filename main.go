package main

import (
	"context"
	"time"
	"todo-cli/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://todo:todo@localhost:27017"))

	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	col := client.Database("public").Collection("todos")

	utils.InsertTodo("Test1", col, ctx)
	utils.InsertTodo("Test2", col, ctx)
	utils.InsertTodo("Test3", col, ctx)
	utils.InsertTodo("Test4", col, ctx)

	utils.UpdateWithDescription("Test2", "In progress", col, ctx)

	utils.Delete("Test3", col, ctx)
}
