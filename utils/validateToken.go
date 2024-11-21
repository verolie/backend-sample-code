package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/code-sample/model/modelDatabase"
	"gorm.io/gorm"
)

func ValidateToken(token string) (modelDatabase.Users, error) {
	db:= SetDatabase()
	
	if (strings.HasPrefix(token, "Bearer ")) {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	
	if token == "" {
		return modelDatabase.Users{}, errors.New("token is null or undefined")
	}

	userId, err := getToken(token, db)
	if err != nil {
		return modelDatabase.Users{}, err
	}

	if userId == 0 {
		return modelDatabase.Users{}, errors.New("invalid token")
	}

	var user modelDatabase.Users
	result := db.Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return modelDatabase.Users{}, errors.New("user not found")
		}
		return modelDatabase.Users{}, fmt.Errorf("error fetching user data: %v", result.Error)
	}

	return user, nil
}

func getToken(token string, db *gorm.DB) (int, error) {
	var data modelDatabase.Tokens

	result := db.Where("token = ?", token).Order("created_at desc").First(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, fmt.Errorf("cannot find data: %v", result.Error)
	}

	return data.UserTokenID, nil
}
