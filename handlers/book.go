package handlers

import "github.com/gofiber/fiber/v2"

type BookHandler interface {
	GetAllBook(c *fiber.Ctx) error
}
