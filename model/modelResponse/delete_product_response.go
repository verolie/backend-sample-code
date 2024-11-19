package modelResponse

type DeleteProductResponse struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
}