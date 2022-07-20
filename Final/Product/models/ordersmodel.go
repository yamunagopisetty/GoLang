package models

type Orders struct {
	ID         string `json:"id" gorm:"primaryKey"`
	OrderID    string `json:"orderid"`
	ProductID  string `json:"productid"`
	CustomerID string `json:"customerid"`
}
