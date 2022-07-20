package productmodels

type Products struct{

	ID      int      `json:"id"`
	Name   string     `json:"name"`
	Number  uint       `json:"number"`
	Category string     `json:"category"`
	Description string    `json:"description"`

}
