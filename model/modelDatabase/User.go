package modelDatabase

import "time"

type User struct {
	UserID 		int `gorm:"primary_key;auto_increment;not_null"`
	Username    string
	Password    string
	CreateAt    time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}