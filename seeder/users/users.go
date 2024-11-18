package seed

import (
	"fmt"
	"time"

	"github.com/code-sample/migrate"
	"github.com/code-sample/utils"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	db := utils.SetDatabase()

	users := []migrate.Users{
		{
			Username: "admin",
			Password: hashPassword("admin123"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			fmt.Println("Error seeding user:", err)
		} else {
			fmt.Println("User seeded:", user.Username)
		}
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return ""
	}
	return string(hashedPassword)
}
