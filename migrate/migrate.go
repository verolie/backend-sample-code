package migrate

import (
	"fmt"
	"time"

	"github.com/code-sample/utils"
)

func Init() {
	db := utils.SetDatabase()
	db.Migrator().CreateTable(&Users{})
	db.Migrator().CreateTable(&Tokens{})
	db.Migrator().CreateTable(&StockProducts{})

	sqlDB, _ := db.DB()
	sqlDB.Close()

	fmt.Println("success migrate")
}

type Users struct {
	UserID    int       `gorm:"primaryKey;autoIncrement;notNull"`
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Tokens struct {
	TokenID   int       `gorm:"primaryKey;autoIncrement;notNull"`
	UserTokenID    int       `gorm:"not null"` 
	Token     string
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	User	  Users `gorm:"foreignKey:UserTokenID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type StockProducts struct {
	StockProductID int       `gorm:"primaryKey;autoIncrement;notNull"`
	ProductName    string
	Quantity       int
	Status         string       `gorm:"column:status"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	CreatedBy      int       `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	UpdatedBy      int       `gorm:"not null"` 

	CreatedByUser Users `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedByUser Users `gorm:"foreignKey:UpdatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}