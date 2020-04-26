package main

import (
	"github.com/DamilolaAlao/go-rest-fiber/book"
	"github.com/DamilolaAlao/go-rest-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx)  {
	c.Send("Hello, World!")
}

func initDatabase()  {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
}

func setupRoutes(app *fiber.App)  {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main()  {
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(3000)
}