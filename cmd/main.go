package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"Id"`
	Manufactured string  `json:"Manufactured"`
	Name         string  `json:"Name"`
	Price        float32 `json:"Price"`
}

// array of Products
var products = []Product{
	{Id: 1, Manufactured: "Apple", Name: "iPhone 13 mini", Price: 599.99},
	{Id: 2, Manufactured: "Apple", Name: "iPhone 14", Price: 700.99},
	{Id: 3, Manufactured: "Apple", Name: "iPhone 12", Price: 499.99},
	{Id: 4, Manufactured: "Samsung", Name: "Galaxy S23", Price: 599.99},
}

// here is the response to request of all products
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// function for POST requests that adds a new product to a list
func postProduct(c *gin.Context) {
	var newProduct Product
	
	if err := c.BindJSON(&newProduct); err != nil {
		log.Printf("%s\n", err)
		return
	}
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductsByBrand(c *gin.Context) {
	brand := c.Param("Manufactured")
	for _, v := range products {
		if v.Manufactured == brand {
			c.IndentedJSON(http.StatusOK, v)
		}
	}
}

func main() {
	router := gin.Default()

	//here we manage the GET request for a specific endpoint
	router.GET("/products", getProducts)
	router.GET("/products/:Manufactured", getProductsByBrand)

	//here we manage the POST request to add a new product
	router.POST("/products", postProduct)

	//here we run a server on a local device with localhost 8080 address
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
