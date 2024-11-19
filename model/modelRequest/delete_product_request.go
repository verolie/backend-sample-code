package modelRequest

type DeleteProductRequest struct {
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
}