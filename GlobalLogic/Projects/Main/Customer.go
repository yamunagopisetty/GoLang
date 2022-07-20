package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customers struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Number  uint   `json:"number"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var CustomersDetails = []Customers{
	{ID: "1", Name: "yamuna", Number: 244, Email: "yamuna@gmail.com", Address: "gooty"},
	{ID: "2", Name: "sai", Number: 245, Email: "sai@gmail.com", Address: "Anantapur"},
	{ID: "3", Name: "teja", Number: 246, Email: "teja@gmail.com", Address: "dhone"},
}

func getCustomers(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, CustomersDetails)
}

func createCustomers(Context *gin.Context) {
	var newCustomer Customers

	if err := Context.BindJSON(&newCustomer); err != nil {
		//Contex.AbortWithError(http.StatusBadReques t,er)
		return
	}

	CustomersDetails = append(CustomersDetails, newCustomer)

	Context.IndentedJSON(http.StatusCreated, newCustomer)

}

func getbyCustomers(Context *gin.Context) {
	id := Context.Param("id")
	Products, err := getCustomerById(id)

	if err != nil {
		Context.IndentedJSON(http.StatusNotFound, gin.H{"message ": "Products not found"})
		return
	}
	Context.IndentedJSON(http.StatusOK, Products)
}
func getCustomerById(id string) (*Customers, error) {
	for i, p := range CustomersDetails {
		if p.ID == id {
			return &CustomersDetails[i], nil
		}
	}
	return nil, errors.New("Products not found")
}

func main() {

	c := gin.Default()

	c.GET("/get", getCustomers)
	c.GET("/getBy/:id", getbyCustomers)
	//c.PATCH("/get/:id", updateOrder)
	c.POST("/create", createCustomers)
	c.Run("localhost:9090")
}
