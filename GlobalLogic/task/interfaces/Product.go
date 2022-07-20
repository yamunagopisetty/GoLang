package interfaces

import "task/models"

type IProduct interface {
	Get() (*models.Products, error)
}
