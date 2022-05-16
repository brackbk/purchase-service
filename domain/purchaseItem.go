package domain

import (
	"gorm.io/gorm"
)

type PurchaseItem struct {
	gorm.Model
	PurchaseId int     `json:"purchase_id" gorm:"type:int"`
	Status     bool    `json:"status" gorm:"type:boolean; default:true"`
	ProductId  int     `json:"product_id" gorm:"type:int"`
	Quantity   int     `json:"quantity" gorm:"type:int"`
	UnitPrice  float64 `json:"unit" gorm:"type:float"`
	TotalPrice float64 `json:"totalprice" gorm:"type:float"`
	CurrencyId int     `json:"currency_id"`
	Purchase   Purchase
}
