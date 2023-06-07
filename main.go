package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Manufactured string  `json:"made-by"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
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

// this func returns all products whose manufacturer is Apple
func getAppleProducts(c *gin.Context) {
	iPhones := func() []Product {
		arr := make([]Product, 0)
		for _, v := range products {
			if v.Manufactured == "Apple" {
				arr = append(arr, v)
			}
		}
		return arr
	}
	res := iPhones()
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	router := gin.Default()

	//here we manage the GET request for a specific endpoint
	router.GET("/products", getProducts)
	router.GET("/products/apple", getAppleProducts)

	//here we run a server on a local device with localhost 8080 address
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
