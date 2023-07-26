package main

import (
	"context"
	"fmt"
	"go-course-ep3/models"
	"go-course-ep3/repositories"
	"time"

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

	bookRepo := repositories.NewBookDBRepository(db, "book_stock")

	// bookRepo.Create(models.RepoBookModel{
	// 	BookID: uuid.New().String(),
	// 	Title:  "GG",
	// 	Price:  100,
	// 	Stock:  20,
	// })

	res, _ := bookRepo.GetAll()
	fmt.Println(res)

	// bookRepo.Delete("f1b7065e-c1a7-4d07-98eb-5c4bae114452")

	bookRepo.Update("0613e971-fa7b-4a16-8032-fbdef6b3d054", models.RepoBookUpdateModel{
		// Title: "GGX",
		Price: 149,
		Stock: 300,
	})

	res, _ = bookRepo.GetAll()
	fmt.Println(res)

	// res2, _ := bookRepo.GetById("0613e971-fa7b-4a16-8032-fbdef6b3d054")
	// fmt.Println(res2)
}
