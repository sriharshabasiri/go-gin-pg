package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sriharshabasiri/go-gin-pg/utils"
	"log"
	"net/http"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {

	properyManager := utils.PropertyManager{}

	database, _ := properyManager.GetProperty("database")
	dbname, _ := properyManager.GetProperty("dbname")
	dbuser, _ := properyManager.GetProperty("dbuser")
	dbpassword, _ := properyManager.GetProperty("dbpassword")
	host, _ := properyManager.GetProperty("host")
	port, _ := properyManager.GetProperty("port")
	sslmode, _ := properyManager.GetProperty("sslmode")

	app_port, _ := properyManager.GetProperty("app_port")

	fmt.Println("database is ** ", database)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, dbuser, dbpassword, dbname, sslmode)

	// Replace the database connection parameters with your own values
	db, err := sql.Open(database, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Gin router
	r := gin.Default()

	// Define the "GET /products" endpoint for retrieving all products
	r.GET("/products", func(c *gin.Context) {

		var product Product
		var products []map[string]interface{}

		rows, err := db.Query("SELECT * FROM product where id>$1", 0)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No Products available"})
			return
		}
		// iterate over the rows
		for rows.Next() {
			err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
			if err != nil {
				panic(err)
			}
			product := map[string]interface{}{
				"ID":          product.ID,
				"Name":        product.Name,
				"Description": product.Description,
				"Price":       product.Price,
			}

			products = append(products, product)
		}
		// Return the product as JSON
		c.JSON(http.StatusOK, products)
	})

	// Define the "GET /products/:id" endpoint for retrieving a product by ID
	r.GET("/products/:id", func(c *gin.Context) {

		// Get the ID parameter from the request URL
		id := c.Param("id")

		// Query the database for the product with the specified ID
		var product Product
		err := db.QueryRow("SELECT * FROM product WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Return the product as JSON
		c.JSON(http.StatusOK, product)
	})

	// Define the "POST /products" endpoint for creating a new product
	r.POST("/products", func(c *gin.Context) {
		// Bind the request body to a Product struct
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(product.Name)
		fmt.Println(product.Description)
		fmt.Println(product.Price)

		// Insert the new product into the database
		var id int
		err := db.QueryRow("INSERT INTO product(product_name, product_desc, price) VALUES($1, $2, $3) RETURNING id", product.Name, product.Description, product.Price).Scan(&id)
		if err != nil {
			fmt.Printf("Error: %+v", errors.New("Exception"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		// Set the ID of the new product and return it as JSON
		product.ID = id
		c.JSON(http.StatusCreated, product)
	})

	// Define the "DELETE /products/:id" endpoint for deleting a product by ID
	r.DELETE("/products/:id", func(c *gin.Context) {
		// Get the ID parameter from the request URL
		id := c.Param("id")

		// Delete the product with the specified ID from the database
		result, err := db.Exec("DELETE FROM product WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		// Check if the product was actually deleted
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Return a success message
		c.JSON(http.StatusOK, gin.H{"error": "Product deleted"})

	})

	r.Run(":" + app_port)
}
