package handlers

import (
	"go-course-ep3/models"
	"go-course-ep3/services"

	"github.com/gofiber/fiber/v2"
)

type bookHandler struct {
	bookSrv services.BookService
}

func NewBookHandler(
	bookSrv services.BookService,
) bookHandler {
	// ) BookHandler {
	return bookHandler{
		bookSrv,
	}
}

func (h bookHandler) GetAllBook(c *fiber.Ctx) error {

	res, err := h.bookSrv.GetAllBook()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get all success",
		"data":    res,
	})
}

func (h bookHandler) GetBookByID(c *fiber.Ctx) error {

	id := c.Params("id")

	res, err := h.bookSrv.GetBookByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get by id success",
		"data":    res,
	})
}

func (h bookHandler) CreateBook(c *fiber.Ctx) error {

	body := models.HandBookModel{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.bookSrv.CreateBook(models.SrvBookModel(body))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "create success",
	})
}

func (h bookHandler) UpdateBook(c *fiber.Ctx) error {

	id := c.Params("id")
	body := models.HandBookUpdateModel{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.bookSrv.UpdateBook(id, models.SrvBookUpdateModel(body))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "update success",
	})
}

func (h bookHandler) DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.bookSrv.DeleteBook(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete success",
	})
}
