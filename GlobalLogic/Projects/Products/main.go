package main

import (
	"error"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Number      uint   `json:"number"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

var ProductsDetails = []Products{
	{ID: 1, Name: "mi", Number: 675, Category: "asd", Description: "communication"},
	{ID: 2, Name: "vivo", Number: 676, Category: "mobile", Description: "communication"},
	{ID: 3, Name: "sam", Number: 677, Category: "lab", Description: "communication"},
}

func getProduct(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, ProductsDetails)
}

func createProduct(Context *gin.Context) {
	var newProducts Products

	if err := Context.BindJSON(&newProducts); err != nil {
		return
	}

	ProductsDetails = append(ProductsDetails, newProducts)

	Context.IndentedJSON(http.StatusCreated, newProducts)

}
func getProductById(id string) (*Products, error.ID) {
	for i, j := range ProductsDetails {
		if j.ID == id {
			return &ProductsDetails[i], nil
		}
		return nil, error.New("Products not found")
	}
}

func getbyProduct(Context *gin.Context) {
	id := Context.param("id")
	Products, err := getProductById(id)

	if err != nil {
		Context.IndentedJSON(http.StatusNotFound, gin.H{"message ": "Products not found"})
		return
	}
	Context.IndentedJSON(http.StatusOK, Products)
}

func main() {
	c := gin.Default()

	c.GET("/get", getProduct)
	c.GET("/getBy/:id", getbyProduct)
	c.POST("/create", createProduct)
	c.Run("localhost:9090")
}
