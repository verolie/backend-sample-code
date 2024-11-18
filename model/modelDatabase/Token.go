package modelDatabase

import "time"

type Tokens struct {
	TokenID     int `gorm:"primaryKey;autoIncrement;notNull"`
	UserTokenID int `gorm:"not null"`
	Token       string
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
