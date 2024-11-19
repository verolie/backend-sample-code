package modelResponse

import "time"

type ProductResponse struct {
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}