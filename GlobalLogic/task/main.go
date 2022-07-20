package main

import (
	"task/database"
	"task/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db := &database.ProductData{}
	ph := &handler.ProductHandler{IProduct: db}

	r := gin.Default()

	r.GET("/get", ph.Gethand())
	r.POST("/create",ph.createhand())
	r.Run()

}
