package fiberHelpers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type RouterConfig struct {
	DB     *sql.DB
	Router fiber.Router
}

// Paths below are created to make it easy to add specific validation per EP as needed & make engineers' lives easier

func (c *RouterConfig) GET(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Get(path, handler)
}

func (c *RouterConfig) POST(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Post(path, handler)
}

func (c *RouterConfig) PUT(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Put(path, handler)
}

func (c *RouterConfig) DELETE(path string, handler fiber.Handler) fiber.Router {
	return c.Router.Delete(path, handler)
}

func RecordNotFound(ctx *fiber.Ctx, err error) error {
	if err.Error() == "sql: no rows in result set" {
		return ctx.Status(fiber.StatusNotFound).JSON(ResourceNotFound)
	} else {
		return ctx.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
	}
}