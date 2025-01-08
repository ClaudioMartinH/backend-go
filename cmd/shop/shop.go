package shop

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"
	// "strconv"
	"fmt"

	database "github.com/ClaudioMartinH/backend-go/cmd/database"
	types "github.com/ClaudioMartinH/backend-go/cmd/models"
	"github.com/gofiber/fiber/v2"
)

var products []types.Product

type cart struct {
	products   []types.Product
	totalPrice float64
}

var Cart = &cart{}

func (c *cart) AddProduct(product *types.Product) {
	c.products = append(c.products, *product)
	c.totalPrice += product.Price
	fmt.Printf("%s was added to cart: %+v\n", product.Name, c.products)
}

func (c *cart) RemoveProduct(id string) {
	for i, p := range c.products {
		if p.Id == id {
			c.products = append(c.products[:i], c.products[i+1:]...)
			c.totalPrice -= p.Price
			break
		}
	}

	// // Update the total price
	for _, p := range c.products {
		c.totalPrice -= p.Price
	}
}

func (c *cart) GetProducts() []types.Product {
	return c.products
}

func HandleAddToCart(c *fiber.Ctx) error {
	productId := c.Params("id")
	product, err := database.FindProductById(productId)
	if err != nil {
		return err
	}
	Cart.AddProduct(product)
	// Update the total price
	// for _, p := range Cart.products {
	// 	Cart.totalPrice += p.Price
	// }
	return c.JSON(Cart.totalPrice)
}

func HandleRemoveFromCart(c *fiber.Ctx) error {
	productId := c.Params("id")
	Cart.RemoveProduct(productId)
	// Update the total price
	for _, p := range Cart.products {
		Cart.totalPrice -= p.Price
	}
	return c.JSON(Cart.products)
}

func HandleGetCart(c *fiber.Ctx) error {
	Cart.totalPrice = 0
	for _, p := range products {
		Cart.totalPrice += p.Price
	}
	return c.JSON(fiber.Map{
		"products":    Cart.GetProducts(),
		"total_price": Cart.totalPrice,
	})
}

func HandleCheckout(c *fiber.Ctx) error {
	Cart.totalPrice = 0
	for _, p := range products {
		Cart.totalPrice += p.Price
	}
	return c.JSON(fiber.Map{
		"message":     "Checkout successful!",
		"total_price": Cart.totalPrice,
	})
}

func ViewCart(c *fiber.Ctx) error {
	for _, p := range products {
		Cart.totalPrice += p.Price
	}
	cartData := types.CartData{
		Products:   Cart.GetProducts(),
		TotalPrice: Cart.totalPrice,
	}
	return c.JSON(cartData)
}

func HandleClearCart(c *fiber.Ctx) error {
	Cart.products = []types.Product{}
	Cart.totalPrice = 0
	return c.JSON(fiber.Map{
		"message": "Cart cleared!",
	})
}
