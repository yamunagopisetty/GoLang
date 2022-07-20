package handler

import (
	"Product/interfaces"
	"Product/models"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsHandler struct {
	IProducts interfaces.IProducts
}

func (p *ProductsHandler) GetHand() func(*gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProducts.Get()
		c.JSON(http.StatusOK, product)
	}
}

func (p *ProductsHandler) CreateHand() func(*gin.Context) {
	return func(c *gin.Context) {
		//_, err := ioutil.ReadAll(c.Request.Body)

		product := &models.Products{}
		err := json.NewDecoder(c.Request.Body).Decode(product)
		//err = json.Unmarshal(buf, err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}

		id, _ := p.IProducts.Create(product)

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		c.Abort()

	}
}

func (p *ProductsHandler) DeleteHand() func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.Params.ByName("name")

		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			c.Abort()
			return
		}
		result, _ := p.IProducts.Delete(name)
		//abc := ProductDb.Products
		//if result.(string) == name {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
			//"Rem":     abc,
		})
		c.Abort()
		//}
	}
}

func (p *ProductsHandler) GetByNameHand() func(*gin.Context) {
	return func(c *gin.Context) {

		name := c.Params.ByName("name")

		result, _ := p.IProducts.GetByName(name)

		c.JSON(http.StatusOK, result)

	}
}

func (p *ProductsHandler) UpdateHand() func(*gin.Context) {
	return func(c *gin.Context) {
		product := &models.Products{}
		err := json.NewDecoder(c.Request.Body).Decode(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		name := c.Params.ByName("name")
		result, _ := p.IProducts.Update(name, product)
		c.JSON(http.StatusOK, result)

	}
}
