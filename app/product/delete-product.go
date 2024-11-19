package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProduct(c *gin.Context) {
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

	responseData, err := deleteProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to delete product data", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessMessage(c, responseData, "Product data deleted successfully")
}

func deleteProductProcess(user modelDatabase.User, c *gin.Context) (gin.H, error) {
	db := utils.SetDatabase()

	productID := c.Param("id")
	if productID == "" {
		return nil, errors.New("product ID is required")
	}

	var product modelDatabase.StockProducts
	if err := db.Where("stock_product_id = ?", productID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if product.Status == "active" {
		product.Status = "inactive"
		product.UpdatedAt = time.Now()
		product.UpdatedBy = user.UserID
		if err := db.Save(&product).Error; err != nil {
			return nil, err
		}
	} else if product.Status == "draft" {
		if err := db.Delete(&product).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid status provided")
	}

	result := gin.H{
		"product_id":   product.StockProductID,
		"product_name": product.ProductName,
		"status":       product.Status,
		"action_by":    user.UserID,
	}

	return result, nil
}
