package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Orders struct {
	OrderID    string `json:"orderid"`
	ProductID  string `json:"product id"`
	CustomerID string `json:"customer id"`
}

var OrdersDetails = []Orders{
	{OrderID: "1", ProductID: "56356", CustomerID: "244"},
	{OrderID: "2", ProductID: "267747", CustomerID: "245"},
	{OrderID: "3", ProductID: "45662", CustomerID: "246"},
}

func getOrders(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, OrdersDetails)
}

func createOrders(Context *gin.Context) {
	var newOrder Orders

	if err := Context.BindJSON(&newOrder); err != nil {
		//Contex.AbortWithError(http.StatusBadReques t,er)
		return
	}

	OrdersDetails = append(OrdersDetails, newOrder)

	Context.IndentedJSON(http.StatusCreated, newOrder)

}

/*func getbyOrders(Context *gin.Context) {
	id := Context.Param("id")
	Orders, err := getOrdersById(id)

	if err != nil {
		Context.IndentedJSON(http.StatusNotFound, gin.H{"message ": "Products not found"})
		return
	}
	Context.IndentedJSON(http.StatusOK, Orders)
}

/*func getOrdersById(id string) (*Orders, error) {
	for i, p := range OrdersDetails {
		if p.ID == id {
			return &OrdersDetails[i], nil
		}
	}
	return nil, errors.New("Products not found")
}*/

func main() {

	c := gin.Default()

	c.GET("/get", getOrders)
	//c.GET("/getBy/:id", getbyOrders)
	//c.PATCH("/get/:id", updateOrder)
	c.POST("/create", createOrders)
	c.Run("localhost:9090")
}
