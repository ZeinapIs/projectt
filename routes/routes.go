// cmd/routes.go

package routes

import (
	"github.com/ZeinapIs/projectt/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Get All Books
	app.Get("/api/books", handlers.GetAllBooks)

	// Get Book Details
	app.Get("/api/books/:bookID", handlers.GetBookDetails)

	// Mark as Currently Reading
	app.Post("/api/books/:bookID/currently-reading", handlers.MarkAsCurrentlyReading)

	// Mark as Read
	app.Post("/api/books/:bookID/mark-as-read", handlers.MarkAsRead)

	// Mark as Did Not Finish
	app.Post("/api/books/:bookID/did-not-finish", handlers.MarkAsDidNotFinish)

	// Get Currently Reading List
	app.Get("/api/currently-reading", handlers.GetCurrentlyReadingList)

	// Get Read Books List
	app.Get("/api/read-books", handlers.GetReadBooksList)

	// Get Did Not Finish List
	app.Get("/api/did-not-finish-books", handlers.GetDidNotFinishList)

	// Add New Book
	app.Post("/api/books", handlers.AddNewBook)

	// Update Book Details
	app.Put("/api/books/:bookID", handlers.UpdateBookDetails)

	// Delete Book
	app.Delete("/api/books/:bookID", handlers.DeleteBook)

	// Search Books
	app.Get("/api/books/search", handlers.SearchBooks)

	// Mark as To Be Read
	app.Post("/api/books/:bookID/to-be-read", handlers.MarkAsToBeRead)

	// Get To Be Read List
	app.Get("/api/to-be-read", handlers.GetToBeReadList)
}
