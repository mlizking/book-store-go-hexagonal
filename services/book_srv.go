package services

import (
	"errors"
	"go-course-ep3/models"
	"go-course-ep3/repositories"

	"github.com/google/uuid"
)

type bookSrv struct {
	bookRepo repositories.BookRepository
}

func NewBookService(
	bookRepo repositories.BookRepository,
) BookService {
	return bookSrv{
		bookRepo,
	}
}

func (s bookSrv) GetAllBook() (result []models.SrvBookModel, err error) {

	res, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		result = append(result, models.SrvBookModel(v))
	}

	return result, nil
}

func (s bookSrv) GetBookByID(bookId string) (result *models.SrvBookModel, err error) {

	if bookId == "" {
		return nil, errors.New("book_id not found")
	}

	res, err := s.bookRepo.GetById(bookId)
	if err != nil {
		return nil, err
	}

	result = &models.SrvBookModel{
		BookID: res.BookID,
		Title:  res.Title,
		Price:  res.Price,
		Stock:  res.Stock,
	}

	return result, nil
}

func (s bookSrv) CreateBook(payload models.SrvBookModel) (err error) {

	if payload.Title == "" {
		return errors.New("title not found")
	}

	err = s.bookRepo.Create(models.RepoBookModel{
		BookID: uuid.New().String(),
		Title:  payload.Title,
		Price:  payload.Price,
		Stock:  payload.Stock,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s bookSrv) UpdateBook(bookId string, payload models.SrvBookUpdateModel) (err error) {

	if bookId == "" {
		return errors.New("book_id not found")
	}

	err = s.bookRepo.Update(bookId, models.RepoBookUpdateModel(payload))
	if err != nil {
		return err
	}

	return nil
}
func (s bookSrv) DeleteBook(bookId string) (err error) {

	if bookId == "" {
		return errors.New("book_id not found")
	}

	err = s.bookRepo.Delete(bookId)
	if err != nil {
		return err
	}

	return nil
}
