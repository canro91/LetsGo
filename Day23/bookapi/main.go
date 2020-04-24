package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/canro91/30DaysOfGo/Day23/bookapi/database"
	"github.com/canro91/30DaysOfGo/Day23/bookapi/book"
	"github.com/gofiber/fiber"
)

func HelloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func withRoutes(app *fiber.App){
	app.Get("/", HelloWorld)

	app.Get("/api/v1/Book", book.GetAllBooks)
	app.Get("/api/v1/Book/:id", book.GetSingleBook)
	app.Post("/api/v1/Book", book.NewBook)
	app.Put("/api/v1/Book/:id", book.UpdateBook)
	app.Delete("/api/v1/Book/:id", book.DeleteBook)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Fail to connect database")
	}
	fmt.Println("Connection opened to database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migration")
}

func main() {
	app := fiber.New()

	initDatabase()

	withRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
