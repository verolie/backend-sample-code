package login

import (
	"fmt"
	"net/http"

	"github.com/code-sample/model/modelDatabase"
	"github.com/code-sample/model/modelRequest"
	"github.com/code-sample/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostLogin(c *gin.Context) {
	var request modelRequest.LoginRequest
	if err := c.Bind(&request); err != nil {
		utils.ErrorMessage(c, "Failed to bind request data", err.Error(), http.StatusBadRequest)
		return
	}

	if request.Username != "" && request.Password != "" {
		data, token, err := loginEmail(request.Username, request.Password)
		if err != nil {
			utils.ErrorMessage(c, "Failed to login", err.Error(), http.StatusUnauthorized)
			return
		}
		
		response := gin.H{
			"username":  data.Username,
			"token": token,
		}
		utils.SuccessMessage(c, response, "Login successful")
	} else {
		utils.ErrorMessage(c, "Failed to login", "Email and password must not be empty", http.StatusBadRequest)
	}
}

func loginEmail(username string, password string) (*modelDatabase.Users, string, error) {
	user := modelDatabase.Users{}
	db := utils.SetDatabase()
	err := db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", fmt.Errorf("incorrect password")
	}

	token, err := utils.ValidateLoginToken(user.UserID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	return &user, token, nil
}
