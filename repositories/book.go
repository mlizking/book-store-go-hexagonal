package repositories

import "go-course-ep3/models"

type BookRepository interface {
	GetAll() (result []models.RepoBookModel, err error)
	GetById(bookId string) (result *models.RepoBookModel, err error)
	Create(payload models.RepoBookModel) (err error)
	Update(bookId string, payload models.RepoBookUpdateModel) (err error)
	Delete(bookId string) (err error)
}
