package modelRequest

type ProductRequest struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Status      string `json:"status"`
}