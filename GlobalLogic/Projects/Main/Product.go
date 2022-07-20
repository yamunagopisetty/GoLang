package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Products struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Number      uint   `json:"number"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

/*var ProductsDetails = []Products{
	{ID: "1", Name: "mi", Number: 675, Category: "asd", Description: "communication"},
	{ID: "2", Name: "vivo", Number: 676, Category: "mobile", Description: "communication"},
	{ID: "3", Name: "sam", Number: 677, Category: "lab", Description: "communication"},
}

func getProduct(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, ProductsDetails)
}*/
func (c *ProductDB) Create(*Products) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&Products{})
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return contact.ID, nil
}

func createProduct(c *gin.Context) {
	var newProducts Products

	if err := c.BindJSON(&newProducts); err != nil {

		return
	}

	ProductsDetails = append(ProductsDetails, newProducts)

	c.IndentedJSON(http.StatusCreated, newProducts)

}

/*func getbyProduct(Context *gin.Context) {
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
}*/
/*create:=Create(*Products) (interface{}, error){


func (ch *ProductHandler) CreateHand() func(*gin.Context) {
	return func(c *gin.Context) {

		buf, _ := ioutil.ReadAll(c.Request.Body)

		contact := &models.Products{}
		err = json.Unmarshal(buf, contact)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		id, _ := ch.IProducts.Create(contact)

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		glog.Info("Product successfully created:", id)
		c.Abort()
	}

}

func (c *ProductDB) create{
	c.Client.(*gorm.DB).AutoMigrate(&Products{})
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return contact.ID, nil
}


}*/

func main() {

	c := gin.Default()

	//c.GET("/get", getProduct)
	r.POST("/Product/Create", CreateHand())
	//c.GET("/getBy/:id", getbyProduct)
	//c.PATCH("/update/:id", updateProduct)
	//c.POST("/create", createProduct)
	c.Run("localhost:9090")
}
