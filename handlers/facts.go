package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Db.Find(&books)
	return c.JSON(books)
}

func GetBookDetails(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	var book models.Book
	result := database.DB.Db.First(&book, bookID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func MarkAsCurrentlyReading(c *fiber.Ctx) error {
	return markBookStatus(c, "currently reading")
}

func MarkAsRead(c *fiber.Ctx) error {
	return markBookStatus(c, "read")
}

func MarkAsDidNotFinish(c *fiber.Ctx) error {
	return markBookStatus(c, "did not finish")
}

func MarkAsToBeRead(c *fiber.Ctx) error {
	return markBookStatus(c, "to read")
}

func markBookStatus(c *fiber.Ctx, status string) error {
	// Get bookID from params
	bookID := c.Params("bookID")

	// Retrieve the book from the database
	var book models.Book
	result := database.DB.Db.First(&book, bookID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Update the book status
	book.Status = status
	database.DB.Db.Save(&book)

	// Return the updated book
	return c.JSON(book)
}

func getBooksByStatus(c *fiber.Ctx, status string) error {
	var books []models.Book
	database.DB.Db.Where("status = ?", status).Find(&books)
	return c.JSON(books)
}

func GetCurrentlyReadingList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "currently reading")
}

func GetReadBooksList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "read")
}

func GetDidNotFinishList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "did not finish")
}

func GetToBeReadList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "to read")
}

func AddNewBook(c *fiber.Ctx) error {
	var newBook models.Book

	// Read the request body
	bodyBytes := c.Body()

	// Unmarshal the request body into the Book struct
	if err := json.Unmarshal(bodyBytes, &newBook); err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Insert the new book into the database
	database.DB.Db.Create(&newBook)

	return c.JSON(newBook)
}
func UpdateBookDetails(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	result := database.DB.Db.Model(&models.Book{}).Where("id = ?", bookID).Updates(&updatedBook)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	c.SendStatus(fiber.StatusNoContent)
	return c.JSON(updatedBook)
}

func DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	result := database.DB.Db.Delete(&models.Book{}, bookID)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	c.SendStatus(fiber.StatusNoContent)
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book was deleted"})
}

func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	var searchResults []models.Book
	result := database.DB.Db.Where("SELECT * FROM books WHERE Title ILIKE ? OR Author ILIKE ?", "%"+query+"%", "%"+query+"%").Scan(&searchResults)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(searchResults)
}
