package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name" gorm:"text;not null;default:null"`
	Price     float32   `json:"price" gorm:"decimal;not null;default:null"`
}

type Order struct {
	ID           uint           `json:"id" gorm:"primary_key;"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
	CustomerName string         `json:"customer_name" gorm:"text;not null;default:null"`
	Items        []ProductItem  `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductItem struct {
	ID        uint      `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UnitPrice float32   `json:"price" gorm:"decimal;not null;default:null"`
	Quantity  int16     `json:"quantity" gorm:"integer;not null;default:null"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ProductID"`
	OrderID   uint      `json:"order_id"`
}
