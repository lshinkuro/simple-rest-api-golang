package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lshinkuro/go-fiber-tutorial/api/book"
	"github.com/lshinkuro/go-fiber-tutorial/api/santri"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")

	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book/*", book.NewBook)
	v1.Put("/book/:id", book.UpdateBook)
	v1.Delete("/book/:id", book.DeleteBook)

	v1.Get("/santri", santri.GetUsers)
	v1.Get("/santri/:id", santri.GetUserById)
	v1.Post("/santri/*", santri.NewUser)
	v1.Put("/santri/:id", santri.UpdateUser)

}
