package product

import (
	"fmt"
	"net/http"
	"time"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/model/modelRequest"
	"github.com/code-sample/model/modelResponse"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
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

    responseData, err := createProductProcess(user, c)
	if err != nil {
		utils.ErrorMessage(c, "Failed to process product data", err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SuccessMessage(c, responseData, "Product data created successfully")
}

func createProductProcess(user modelDatabase.User, c *gin.Context) (modelResponse.ProductResponse, error) {
	var reqProduct modelRequest.CreateProductRequest
    db := utils.SetDatabase()

	if err := c.Bind(&reqProduct); err != nil {
		return modelResponse.ProductResponse{}, err
	}

    if (reqProduct.Status == "active" || reqProduct.Status == "draft") {

    	product :=  modelDatabase.StockProducts{
            ProductName: reqProduct.ProductName,
            Quantity:    reqProduct.Quantity,
            Status:      reqProduct.Status,
            CreatedAt:   time.Now(),
            CreatedBy:   user.UserID, 
            UpdatedAt:   time.Now(), 
            UpdatedBy:   user.UserID,
        }

    	if err := db.Create(&product).Error; err != nil {
    		return modelResponse.ProductResponse{}, err
    	}

    	response := modelResponse.ProductResponse {
            ProductID:   product.StockProductID,
            ProductName: product.ProductName,
            Quantity:    product.Quantity,
            Status:      product.Status,
            CreatedAt:   product.CreatedAt,
            CreatedBy:   product.CreatedBy,
            UpdatedAt:   product.UpdatedAt,
            UpdatedBy:   product.UpdatedBy,
        }
    

	    return response, nil
    } else{
        return modelResponse.ProductResponse{}, fmt.Errorf("error: invalid status")
    }
}
