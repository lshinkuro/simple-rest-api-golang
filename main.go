package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lshinkuro/go-fiber-tutorial/book"
	"github.com/lshinkuro/go-fiber-tutorial/database"
	"github.com/lshinkuro/go-fiber-tutorial/middleware"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Succesfully connect")
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book/*", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0", "1.1.1.1/30"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	initDatabase()

	middleware.MiddlewareApi(app)
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

	defer database.DBConn.Close()
}
