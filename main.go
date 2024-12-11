package main

import (
	"log"

	database "github.com/ClaudioMartinH/backend-go/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	app := fiber.New()
	// Middleware
	app.Use(logger.New())
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Routes Grouped
	userGroup := app.Group("/users")
	userGroup.Get("/:id", database.HandleUser)
	userGroup.Post("", database.HandleCreateUser)

	// Middleware
	app.Use(requestid.New())

	// Routes Grouped
	productsGroup := app.Group("/products")
	productsGroup.Get("/:id", database.HandleProduct)
	productsGroup.Post("", database.HandleCreateProduct)

	// Listen like in Express
	app.Listen(":3000")
}
