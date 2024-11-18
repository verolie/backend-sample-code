package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	CONST_PG_HOST     string = "localhost"
	CONST_PG_PORT     string = "3333"
	CONST_PG_USER     string = "admin"
	CONST_PG_PASSWORD string = "pass123"
	CONST_PG_DBNAME   string = "projectLaravel"
)

func SetDatabase() *gorm.DB {

	host := checkConditionConst("PG_HOST")
	port := checkConditionConst("PG_PORT")
	user := checkConditionConst("PG_USER")
	password := checkConditionConst("PG_PASSWORD")
	dbname := checkConditionConst("PG_DBNAME")

	// PostgreSQL connection string format
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(dsn)

	// Use the PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}
	return db
}

func checkConditionConst(constValue string) string {
	value := os.Getenv(constValue)
	if value != "" {
		return value
	} else {
		// Replace MySQL-specific constants with PostgreSQL-specific ones
		switch constValue {
		case "PG_HOST":
			return CONST_PG_HOST
		case "PG_PORT":
			return CONST_PG_PORT
		case "PG_USER":
			return CONST_PG_USER
		case "PG_PASSWORD":
			return CONST_PG_PASSWORD
		case "PG_DBNAME":
			return CONST_PG_DBNAME
		}
	}

	return value
}
