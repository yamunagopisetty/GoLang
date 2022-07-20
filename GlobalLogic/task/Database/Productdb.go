package database

import (
	"fmt"
	"task/models"
)

type ProductData struct {
}

func (p *ProductData) Get() (*models.Products, error) {
	m := &models.Products{ID: 1, Name: "Mi", Number: 23455, Category: "mobile", Description: "communication"}
	ErrorType := fmt.Errorf("there is Error in Products Database")
	return m, ErrorType
}
