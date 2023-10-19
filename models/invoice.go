package models


type Invoice struct {
	Id_product 	int 	`json:"id_product"`
	Product 	string 	`json:"product"`
	Quantity 	int 	`json:"quantity"`
	Price 		float64 `json:"price"`
	Tax_rate 	float64 	`json:"tax_rate"`
	Discount_rate float64 `json:"discount_rate"`
	Currency string `json:"currency"`
}
