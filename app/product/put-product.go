package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/model/modelRequest"
	"github.com/code-sample/model/modelResponse"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutProduct(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ErrorMessage(c, "Authorization header is required", "missing authorization header", http.StatusUnauthorized)
		return
	}

	user, err := utils.ValidateToken(authHeader)
	if err != nil {
		utils.ErrorMessage(c, "Invalid or missing token", err.Error(), http.StatusUnauthorized)
		return
	}

	responseData, err := putProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to update product data", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessMessage(c, responseData, "Product data updated successfully")
}

func putProductProcess(user modelDatabase.Users, c *gin.Context) (modelResponse.ProductResponse, error) {
	db := utils.SetDatabase()

	productID := c.Param("id")
	if productID == "" {
		return modelResponse.ProductResponse{}, errors.New("product ID is required")
	}

	var reqProduct modelRequest.CreateProductRequest
	if err := c.Bind(&reqProduct); err != nil {
		return modelResponse.ProductResponse{}, err
	}

	if err := validateStatus(reqProduct.Status); err != nil {
		return modelResponse.ProductResponse{}, err
	}

	var product modelDatabase.StockProducts
	if err := db.Where("stock_product_id = ?", productID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return modelResponse.ProductResponse{}, errors.New("product not found")
		}
		return modelResponse.ProductResponse{}, err
	}

	updatedProduct, err := updateProduct(reqProduct, product, user, db)
	if err != nil {
		return modelResponse.ProductResponse{}, err
	}

	response := modelResponse.ProductResponse{
		ProductID:   updatedProduct.StockProductID,
		ProductName: updatedProduct.ProductName,
		Quantity:    updatedProduct.Quantity,
		Status:      updatedProduct.Status,
		CreatedAt:   updatedProduct.CreatedAt,
		CreatedBy:   user.Username,
		UpdatedAt:   updatedProduct.UpdatedAt,
		UpdatedBy:   user.Username,
	}

	return response, nil
}

func validateStatus(status string) error {
	validStatuses := []string{"active", "draft", "inactive"}
	isValidStatus := false

	for _, validStatus := range validStatuses {
		if status == validStatus {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		return errors.New("invalid status")
	}

	return nil
}

func updateProduct(reqProduct modelRequest.CreateProductRequest, product modelDatabase.StockProducts, user modelDatabase.Users, db *gorm.DB) (*modelDatabase.StockProducts, error) {
	if reqProduct.ProductName != "" {
		product.ProductName = reqProduct.ProductName
	}
	if reqProduct.Quantity != 0 {
		product.Quantity = reqProduct.Quantity
	}
	if reqProduct.Status != "" {
		product.Status = reqProduct.Status
	}

	product.UpdatedAt = time.Now()
	product.UpdatedBy = user.UserID

	if err := db.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
