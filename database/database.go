package database

import (
	"database/sql"
	"time"

	"github.com/ClaudioMartinH/backend-go/connection"
	types "github.com/ClaudioMartinH/backend-go/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleCreateProduct(c *fiber.Ctx) error {
	product := types.Product{}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	product.Id = uuid.NewString()

	// Connect to the database
	db, err := connection.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close() // Close the connection after the function finishes

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS godatabase;")
	if err != nil {
		return err
	}

	// Use the database
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Create the products table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id VARCHAR(36) PRIMARY KEY, Name VARCHAR(255), Description VARCHAR(255), Price DECIMAL(10,2));")
	if err != nil {
		return err
	}

	// Prepare and execute the insert statement
	stmt, err := db.Prepare("INSERT INTO products (id, Name, Description, Price) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the prepared statement

	_, err = stmt.Exec(product.Id, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}
func HandleProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// Connect to the database
	db, err := connection.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close() // Close the connection after the function finishes

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS godatabase;")
	if err != nil {
		return err
	}

	// Use the database
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Create the products table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id VARCHAR(36) PRIMARY KEY, Name VARCHAR(255), Description VARCHAR(255), Price DECIMAL(10,2));")
	if err != nil {
		return err
	}

	// Prepare and execute the select statement
	stmt, err := db.Prepare("SELECT * FROM products WHERE id =?;")
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the prepared statement

	var product types.Product
	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	} else if err != nil {
		return err
	}
	return c.JSON(product)
}

func HandleCreateUser(c *fiber.Ctx) error {
	user := types.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()

	// Connect to the database
	db, err := connection.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close() // Close the connection after the function finishes

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS godatabase;")
	if err != nil {
		return err
	}

	// Use the database
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Create the users table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id VARCHAR(36) PRIMARY KEY, firstname VARCHAR(255), lastname VARCHAR(255));")
	if err != nil {
		return err
	}

	// Prepare and execute the insert statement
	stmt, err := db.Prepare("INSERT INTO users (id, firstname, lastname) VALUES (?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the prepared statement

	_, err = stmt.Exec(user.Id, user.Firstname, user.Lastname)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func HandleUser(c *fiber.Ctx) error {

	id := c.Params("id")
	// Connect to the database
	db, err := connection.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close() // Close the connection after the function finishes

	// Use the database
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Create the users table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id VARCHAR(36) PRIMARY KEY, firstname VARCHAR(255), lastname VARCHAR(255));")
	if err != nil {
		return err
	}

	// Prepare and execute the select statement
	stmt, err := db.Prepare("SELECT * FROM users WHERE id =?;")
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the prepared statement
	var user types.User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Firstname, &user.Lastname)
	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	} else if err != nil {
		return err
	}

	return c.JSON(user)

}
