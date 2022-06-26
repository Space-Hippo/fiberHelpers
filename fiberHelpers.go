package fiberHelpers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type RouterConfig struct {
	DB        *sql.DB
	VerRouter fiber.Router
}

// Paths below are created to make it easy to add specific validation per EP as needed & make engineers' lives easier

func (c *RouterConfig) GET(path string, handler fiber.Handler) fiber.Router {
	return c.VerRouter.Get(path, handler)
}

func (c *RouterConfig) POST(path string, handler fiber.Handler) fiber.Router {
	return c.VerRouter.Post(path, handler)
}

func (c *RouterConfig) PUT(path string, handler fiber.Handler) fiber.Router {
	return c.VerRouter.Put(path, handler)
}

func (c *RouterConfig) DELETE(path string, handler fiber.Handler) fiber.Router {
	return c.VerRouter.Delete(path, handler)
}

func RecordNotFound(ctx *fiber.Ctx, err error) error {
	if err.Error() == "sql: no rows in result set" {
		return ctx.Status(fiber.StatusNotFound).JSON("Resource not found")
	} else {
		return ctx.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}
}