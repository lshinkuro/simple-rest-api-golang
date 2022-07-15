package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareApi(app *fiber.App) {
	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("🥇 First handler")
		return c.Next()
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥈 Second handler")
		return c.Next()
	})
}
