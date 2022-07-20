package models

type Products struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Number      uint64 `json:"number"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
