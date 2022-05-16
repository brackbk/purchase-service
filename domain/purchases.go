package domain

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	PaymentStatus string     `json:"payment_staus" gorm:"type:varchar(255)"`
	PaymentType   string     `json:"payment_type" gorm:"type:varchar(255)"`
	InvoiceNumber string     `json:"invoice_number" gorm:"type:varchar(255)"`
	Status        bool       `json:"status" gorm:"type:boolean; default:true"`
	CompanyId     *int       `json:"company_id" gorm:"type:int"`
	ProviderId    *int       `json:"provider_id" gorm:"type:int"`
	LocationId    *int       `json:"location_id" gorm:"type:int"`
	CurrencyId    *int       `json:"currency_id" gorm:"type: int"`
	Total         *float64   `json:"total" gorm:"type:float"`
	PurchaseDate  *time.Time `json:"purchaseDate"`
	PurchaseItem  []PurchaseItem
}
