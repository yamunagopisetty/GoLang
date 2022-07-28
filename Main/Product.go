package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Number      uint   `json:"number"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

var ProductsDetails = []Products{
	{ID: "1", Name: "mi", Number: 675, Category: "asd", Description: "communication"},
	{ID: "2", Name: "vivo", Number: 676, Category: "mobile", Description: "communication"},
	{ID: "3", Name: "sam", Number: 677, Category: "lab", Description: "communication"},
}

func getProduct(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, ProductsDetails)
}

func createProduct(c *gin.Context) {
	var newProducts Products

	if err := c.BindJSON(&newProducts); err != nil {

		return
	}

	ProductsDetails = append(ProductsDetails, newProducts)

	c.IndentedJSON(http.StatusCreated, newProducts)

}

func getbyProduct(Context *gin.Context) {
	id := Context.Param("id")
	Products, err := getProductById(id)

	if err != nil {
		Context.IndentedJSON(http.StatusNotFound, gin.H{"message ": "Products not found"})
		return
	}
	Context.IndentedJSON(http.StatusOK, Products)
}
func getProductById(id string) (*Products, error) {
	for i, j := range ProductsDetails {
		if j.ID == id {
			return &ProductsDetails[i], nil
		}
	}
	return nil, errors.New("Products not found")
}

/*func updateProduct(Context *gin.Context) {
	id := Context.Param("id")
	Products, err := getProductById(id)

	if err != nil {
		Context.IndentedJSON(http.StatusNotFound, gin.H{"message ": "Products not found"})
		return
	}
	Products completed = !Products completed

	Context.IndentedJSON(http.StatusOK, Products)
}*/

/*func deleteProduct(c *gin.Context) {
	//var delProduct Products
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "id cannot be empty",
		})
		c.Abort()
		return

	}
	//result, _ := Products.ProductsDetails.Delete(id)
	//ProductsDetails = delete(Products, delProduct)

}*/

func main() {

	c := gin.Default()

	c.GET("/get", getProduct)
	c.GET("/getBy/:id", getbyProduct)
	//c.PATCH("/update/:id", updateProduct)
	c.POST("/create", createProduct)
	c.Run("localhost:9090")
}
