package main

import (
	"log"

	database "github.com/ClaudioMartinH/backend-go/cmd/database"
	shop "github.com/ClaudioMartinH/backend-go/cmd/shop"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	// Middleware de CORS (cambiar en produccion)
	//  app.Use(cors.New(cors.Config{
	//       AllowOrigins:     "http://localhost:5174", // Origen de tu frontend
	//       AllowMethods:     "GET,POST,HEAD,PUT,PATCH,DELETE",
	//       AllowHeaders:     "Origin, Content-Type, Accept",
	//       AllowCredentials: true, // Si necesitas enviar cookies o autenticación
	//   }))
	app.Use(cors.New())
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Routes Grouped
	userGroup := app.Group("/users")
	userGroup.Get("/all", database.GetAllUsers)
	userGroup.Get("/:id", database.HandleUser)
	userGroup.Post("", database.HandleCreateUser)
	userGroup.Put("/edit/:id", database.HandleEditUser)
	userGroup.Delete("/delete/:id", database.HandleDeleteUser)

	// Middleware
	app.Use(requestid.New())

	// Routes Grouped
	productsGroup := app.Group("/products")
	productsGroup.Get("/all", database.GetAllProducts)
	productsGroup.Get("/:id", database.HandleProduct)
	productsGroup.Post("", database.HandleCreateProduct)
	productsGroup.Put("/edit/:id", database.HandleEditProduct)
	productsGroup.Delete("/delete/:id", database.HandleDeleteProduct)

	// Cart Routes Grouped
	cartGroup := app.Group("/cart")
	cartGroup.Post("/add/:id", shop.HandleAddToCart)
	cartGroup.Post("/remove", shop.HandleRemoveFromCart)
	cartGroup.Get("/view", shop.ViewCart)
	cartGroup.Get("/", shop.HandleGetCart)
	cartGroup.Post("/clear", shop.HandleClearCart)
	cartGroup.Post("/checkout", shop.HandleCheckout)

	// Listen like in Express
	app.Listen(":8080")
}
