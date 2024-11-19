package modelDatabase

import "time"

type StockProducts struct {
	StockProductID int 		`gorm:"primaryKey;autoIncrement;notNull"`
	ProductName    string
	Quantity       int
	Status         string    `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	CreatedBy      int       `gorm:"column:created_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	UpdatedBy      int       `gorm:"column:updated_by"`
}