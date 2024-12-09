package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lshinkuro/go-fiber-tutorial/api/book"
	"github.com/lshinkuro/go-fiber-tutorial/api/database"
	"github.com/lshinkuro/go-fiber-tutorial/api/middleware"
	routes "github.com/lshinkuro/go-fiber-tutorial/api/route"
	"github.com/lshinkuro/go-fiber-tutorial/api/santri"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Succesfully connect")
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{}, &santri.User{})
	fmt.Println("Database Migrated")
}

func Run() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0", "1.1.1.1/30"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	initDatabase()

	middleware.MiddlewareApi(app)
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

	defer database.DBConn.Close()
}
