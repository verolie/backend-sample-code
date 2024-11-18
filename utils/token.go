package utils

import (
	"fmt"
	"time"

	"github.com/code-sample/model/modelDatabase"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var secretKey = []byte("secret-key")

func ValidateToken(userId int) (string, error) {
	token, err := checkAvailableToken(userId)
	if err != nil {
		return "", err
	}

	if token == "" {
		newToken, err := createToken(userId)
		if err != nil {
			return "", err
		}
		return newToken, nil
	}

	return token, nil
}

func checkAvailableToken(userId int) (string, error) {
	db := SetDatabase()

	tokenString, err := getToken(userId, db)
	if err != nil {
		return "", err
	}

	if tokenString != "" {
		if err := verifyToken(tokenString); err == nil {
			return tokenString, nil
		}
		err = deleteToken(userId, db)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

func getToken(userId int, db *gorm.DB) (string, error) {
	var data modelDatabase.Tokens

	result := db.Where("user_token_id = ?", userId).Order("created_at desc").First(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "", nil
		}
			return "", fmt.Errorf("cannot find data: %v", result.Error)
	}

	return data.Token, nil
}

func createToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// Save the new token to the database
	err = saveToken(userId, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func saveToken(userId int, tokenString string) error {
	db := SetDatabase()

	// Create a new token entry
	newToken := modelDatabase.Tokens{
		UserTokenID: userId,
		Token:  tokenString,
	}

	result := db.Create(&newToken)
	if result.Error != nil {
		return fmt.Errorf("cannot save token: %v", result.Error)
	}

	return nil
}

func deleteToken(userId int, db *gorm.DB) error {
	// Delete the expired token for the user
	result := db.Where("user_id = ?", userId).Delete(&modelDatabase.Tokens{})
	if result.Error != nil {
		return fmt.Errorf("cannot delete token: %v", result.Error)
	}
	return nil
}

func verifyToken(tokenString string) error {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
