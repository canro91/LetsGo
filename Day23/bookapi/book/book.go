package book

import (
	"github.com/canro91/30DaysOfGo/Day23/bookapi/database"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
)

type Book struct {
	gorm.Model
	Title string `json:"name"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetAllBooks(c *fiber.Ctx){
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetSingleBook(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book == (Book{}) {
		c.Status(400).Send("Book not found")
	} else {
		c.JSON(book)
	}
}

func NewBook(c *fiber.Ctx){
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx){
	db := database.DBConn
	id := c.Params("id")
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(400).Send(err)
		return
	}
	
	var found Book
	db.Find(&found, id)
	if found == (Book{}) {
		c.Status(400).Send("Book not found")
	} else {
		found.Title = book.Title
		found.Author = book.Author
		found.Rating = book.Rating
		db.Save(&found)
		c.JSON(found)
	}
}

func DeleteBook(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book == (Book{}) {
        c.Status(400).Send("No Book Found with ID")
	} else {
		db.Delete(&book)
		c.Send("Book Successfully deleted")
	}
}