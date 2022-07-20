package main

import (
	"Product/ProductDb"
	"Product/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	pdb := &ProductDb.ProductsDB{}
	ph := &handler.ProductsHandler{IProducts: pdb}

	r := gin.Default()

	r.GET("/get", ph.GetHand())
	r.GET("/getBy/:id", ph.GetByNameHand())
	r.POST("/create", ph.CreateHand())
	r.PUT("/update/:id", ph.UpdateHand())
	r.DELETE("/delete/:id", ph.DeleteHand())
	r.Run()

}
