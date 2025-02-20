package main

import (
	"log"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoInstance struct {
	client 
	db
}
var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:uri " + dbName

type Employee struct {
	ID string
	Name  string
	Salary float64 
	Age float64
}

func Connection() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	context.WithTimeOut(context.Background() 30 * time.Second)
}

func main() {
	if err := Connection(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		
	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")
