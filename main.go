package main

import (
	"context"
	"go-course-ep3/handlers"
	"go-course-ep3/repositories"
	"go-course-ep3/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ep5-course:HlT9NpyD4Vt0HtbK@cluster0.vvx397a.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	// # Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client.Database("BookStore")
}

func main() {

	db := initMongo()

	//init Data Layer
	bookRepo := repositories.NewBookDBRepository(db, "book_stock")

	//init Business Logic Layer
	bookSrv := services.NewBookService(bookRepo)

	//init Presentation Layer
	bookHand := handlers.NewBookHandler(bookSrv)

	//framework routes
	app := fiber.New()
	app.Get("/books", bookHand.GetAllBook)
	app.Get("/book/:id", bookHand.GetBookByID)
	app.Post("/book", bookHand.CreateBook)
	app.Put("/book/:id", bookHand.UpdateBook)
	app.Delete("/book/:id", bookHand.DeleteBook)

	//start server
	app.Listen(":3000")
}
