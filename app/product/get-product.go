package product

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProduct handles the retrieval of products with pagination
func GetProduct(c *gin.Context) {
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

	responseData, err := getProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to retrieve product data", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessMessage(c, responseData, "Product data retrieved successfully")
}

// getProductProcess retrieves products from the database with pagination
func getProductProcess(user modelDatabase.User, c *gin.Context) (gin.H, error) {
	db := utils.SetDatabase()

	productId := c.DefaultQuery("product_id", "")

	if productId != "" {
		product, err := GetProductByID(db, productId)
		if err != nil {
			return nil, err
		}

		// Return the product data
		return gin.H{
			"product": product,
		}, nil
	}

	// Pagination parameters
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return nil, errors.New("invalid page number")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		return nil, errors.New("invalid page size")
	}

	offset := (page - 1) * pageSize

	var products []modelDatabase.StockProducts
	if err := db.Limit(pageSize).Offset(offset).Find(&products).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no products found")
		}
		return nil, err
	}

	var totalProducts int64
	if err := db.Model(&modelDatabase.StockProducts{}).Count(&totalProducts).Error; err != nil {
		return nil, err
	}

	result := gin.H{
		"user": user.UserID,
		"products": products,
		"pagination": gin.H{
			"page":      page,
			"pageSize":  pageSize,
			"total":     totalProducts,
			"totalPage": (totalProducts + int64(pageSize) - 1) / int64(pageSize), // Calculate total pages
		},
	}

	return result, nil
}



func GetProductByID(db *gorm.DB, productId string) (*modelDatabase.StockProducts, error) {
	var product modelDatabase.StockProducts
	if err := db.Where("stock_product_id = ?", productId).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}