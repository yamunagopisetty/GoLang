package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Productshandler struct {
}

func (p *Productshandler) getProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProduct.Get()
		c.JSON(http.StatusOK, *product)
	}
}
