package model

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryID  string  `json:"categoryId"`
}

type Category struct {
	Name string `json:"name"`
}
