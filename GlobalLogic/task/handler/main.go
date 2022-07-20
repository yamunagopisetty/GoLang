package handler

import (
	"net/http"
	"task/interfaces"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	IProduct interfaces.IProduct
}

func (p *ProductHandler) Gethand() func(*gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProduct.Get()
		c.JSON(http.StatusOK, *product)

	}

func (p *ProductHandler) createhand() func(*gin.Context){
	return func(c *gin.Context)

	
}
	
}

}
