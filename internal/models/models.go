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
	TotalPrice   float32        `json:"total_price" gorm:"-"`
	CustomerName string         `json:"customer_name" gorm:"text;not null;default:null"`
	Items        []ProductItem  `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (o *Order) AfterFind(tx *gorm.DB) (err error) {
	for _, item := range o.Items {
		o.TotalPrice += item.UnitPrice * float32(item.Quantity)
	}
	return
}

type ProductItem struct {
	ID        uint      `json:"id" gorm:"primary_key;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UnitPrice float32   `json:"price" gorm:"decimal;not null;default:null"`
	Quantity  int16     `json:"quantity" gorm:"integer;not null;default:null"`
	ProductID uint      `json:"product_id"`
	Product   Product   `json:"product"`
	OrderID   uint      `json:"order_id"`
}
