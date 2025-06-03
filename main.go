package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file", err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	collection = client.Database("go-todo-app").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", addTodos)
	// app.Patch("/api/todos:id", updateTodos)
	// app.Delete("/api/todos/:id", delTodos)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

	fmt.Println("Connected To MongoDB Atlas")
}

func getTodos(c *fiber.Ctx) error {

	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)

}

// func addTodos(c *fiber.Ctx) error {

// }

// func updateTodos(c *fiber.Ctx) error {

// }

// func delTodos(c *fiber.Ctx) error {

// }
