package repositories

import (
	"context"
	"errors"
	"go-course-ep3/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookDBRepo struct {
	db         *mongo.Database
	collection string
}

func NewBookDBRepository(
	db *mongo.Database,
	collection string,
) BookRepository {
	return bookDBRepo{
		db:         db,
		collection: collection,
	}
}

func (r bookDBRepo) GetAll() (result []models.RepoBookModel, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r bookDBRepo) GetById(bookId string) (result *models.RepoBookModel, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.db.Collection(r.collection).FindOne(ctx, bson.D{{Key: "book_id", Value: bookId}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r bookDBRepo) Create(payload models.RepoBookModel) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = r.db.Collection(r.collection).InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (r bookDBRepo) Update(bookId string, payload models.RepoBookUpdateModel) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	setUpdateData := bson.D{{Key: "$set", Value: payload}}

	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"book_id": bookId}, setUpdateData)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("mongo: no documents is match")
	}

	return nil
}

func (r bookDBRepo) Delete(bookId string) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"book_id": bookId})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("not delete documents")
	}

	return nil
}
