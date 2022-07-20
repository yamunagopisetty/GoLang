package productsDb

import (
	"Product/models"
	"fmt"
)

type ProductsDB struct {
}

var Products []models.Products

func (p *ProductsDB) Get() (*[]models.Products, error) {

	ErrorType := fmt.Errorf("there is Error in Products Database")
	return &Products, ErrorType
}

func (p *ProductsDB) Create(c *models.Products) (*[]models.Products, error) {
	//pd := &models.Products{}

	Products = append(Products, *c)
	return &Products, nil
}

func (p *ProductsDB) Delete(s string) (string, error) {

	for index, item := range Products {
		if item.Name == s {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}
	return s, nil
}

func (p *ProductsDB) GetByName(s string) (*[]models.Products, error) {
	var Products1 []models.Products
	for index, item := range Products {
		if item.Name == s {
			Products1 = append(Products1, Products[index])
		}
	}
	return &Products1, nil
}

func (p *ProductsDB) Update(s string, t *models.Products) (*[]models.Products, error) {
	for index, item := range Products {
		if item.Name == s {
			Products[index].Number = t.Number
			Products[index].Category = t.Category
			Products[index].Description = t.Description
		}
	}
	return &Products, nil
}
