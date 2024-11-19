package product

import (
	"errors"
	"net/http"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteProduct handles the deletion of a product
func DeleteProduct(c *gin.Context) {
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

	// Process the deletion of the product
	responseData, err := deleteProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to delete product data", err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a successful response
	utils.SuccessMessage(c, responseData, "Product data deleted successfully")
}

// deleteProductProcess deletes a product from the database
func deleteProductProcess(user modelDatabase.User, c *gin.Context) (gin.H, error) {
	db := utils.SetDatabase()

	// Get the product ID from the URL parameter
	productID := c.Param("id")
	if productID == "" {
		return nil, errors.New("product ID is required")
	}

	// Find the product in the database
	var product modelDatabase.StockProducts
	if err := db.Where("product_id = ?", productID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Delete the product
	if err := db.Delete(&product).Error; err != nil {
		return nil, err
	}

	// Return response data after deletion
	result := gin.H{
		"deleted_product_id": productID,
		"deleted_by":         user.UserID,
	}

	return result, nil
}
