package services

import "go-course-ep3/models"

type BookService interface {
	GetAllBook() (result []models.SrvBookModel, err error)
	GetBookByID(bookId string) (result *models.SrvBookModel, err error)
	CreateBook(payload models.SrvBookModel) (err error)
	UpdateBook(bookId string, payload models.SrvBookUpdateModel) (err error)
	DeleteBook(bookId string) (err error)
}
