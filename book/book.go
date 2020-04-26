package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx)  {
	c.Send("Get Books")
}

func GetBook(c *fiber.Ctx)  {
	c.Send("Get Book")
}

func NewBook(c *fiber.Ctx)  {
	c.Send("Get New Book")
}

func DeleteBook(c *fiber.Ctx)  {
	c.Send("Delete Book")
}