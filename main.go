package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool          `json:"completed"`
	Body      string        `json:"body"`
}

var collection *mongo.Collection

func main() {

	// Loding ENV File
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file", err)
	}

	// Mongo Connection
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

	// Database Connection
	collection = client.Database("go-todo-app").Collection("todos")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://todo-service.default.svc.cluster.local/api",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", addTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", delTodos)

	// App Configuration
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

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}
	return c.JSON(todos)

}

func addTodos(c *fiber.Ctx) error {

	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"Error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(bson.ObjectID)
	return c.Status(201).JSON(todo)

}

func updateTodos(c *fiber.Ctx) error {

	id := c.Params("id")

	objectID, err := bson.ObjectIDFromHex(id)

	fmt.Println(id)
	fmt.Println(objectID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update todo"})
	}

	return c.Status(200).JSON(fiber.Map{"success": "Record Updated Successfully"})

}

func delTodos(c *fiber.Ctx) error {

	id := c.Params("id")

	objectID, err := bson.ObjectIDFromHex(id)

	fmt.Println(id)
	fmt.Println(objectID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo ID"})
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo"})
	}

	return c.Status(200).JSON(fiber.Map{"success": "Record deleted successfully"})
}
