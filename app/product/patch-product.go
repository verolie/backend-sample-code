package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/model/modelRequest"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PatchProduct handles the update of a product
func PatchProduct(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.ErrorMessage(c, "Authorization header is required", "missing authorization header", http.StatusUnauthorized)
		return
	}

	// Validate the token
	user, err := utils.ValidateToken(authHeader)
	if err != nil {
		utils.ErrorMessage(c, "Invalid or missing token", err.Error(), http.StatusUnauthorized)
		return
	}

	// Process the update of the product
	responseData, err := patchProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to update product data", err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a successful response
	utils.SuccessMessage(c, responseData, "Product data updated successfully")
}

// patchProductProcess updates a product in the database
func patchProductProcess(user modelDatabase.User, c *gin.Context) (gin.H, error) {
	db := utils.SetDatabase()

	// Get the product ID from the URL parameter
	productID := c.Param("id")
	if productID == "" {
		return nil, errors.New("product ID is required")
	}

	// Bind the incoming request data to the ProductRequest model
	var reqProduct modelRequest.ProductRequest
	if err := c.Bind(&reqProduct); err != nil {
		return nil, err
	}

	// Find the existing product in the database
	var product modelDatabase.StockProducts
	if err := db.Where("product_id = ?", productID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Update the product fields
	product.ProductName = reqProduct.ProductName
	product.Quantity = reqProduct.Quantity
	product.Status = reqProduct.Status
	product.UpdatedAt = time.Now()
	product.UpdatedBy = user.UserID

	// Save the updated product to the database
	if err := db.Save(&product).Error; err != nil {
		return nil, err
	}

	// Return the updated product information in the response
	result := gin.H{
		"updated_product_id": productID,
		"product_name":       product.ProductName,
		"quantity":           product.Quantity,
		"status":             product.Status,
		"updated_by":         user.UserID,
	}

	return result, nil
}
