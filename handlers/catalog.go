package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type CatalogHandlers interface {
	GetProducts(c *fiber.Ctx) error
}
