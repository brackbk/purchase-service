package purchase

import (
	"time"

	"github.com/brackbk/purchase-service/dto/purchaseItem"
)

type RegisterRequest struct {
	PaymentStatus string                         `json:"paymentStatus"`
	PaymentType   string                         `json:"paymentType"`
	CompanyId     *int                           `json:"companyId" validate:"gt=0"`
	ProviderId    *int                           `json:"providerId" validate:"gt=0"`
	LocationId    *int                           `json:"locationId" validate:"gt=0"`
	CurrencyId    *int                           `json:"currencyId" validate:"gt=0"`
	PurchaseItem  []purchaseItem.RegisterRequest `json:"purchaseitems" validate:"required,min=1,dive,required"`
	PurchaseDate  *time.Time                     `json:"purchaseDate"`
	Total         int                            `json:"total"`
}

func (r RegisterRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "Is required and must be a valid Id "
	switch s {
	case "CompanyId":
		field = "companyId"
		message += "Int"
	case "ProviderId":
		field = "providerId"
		message += "Int"
	case "CurrencyId":
		field = "currencyId"
		message += "Int"
	case "LocationId":
		field = "locationId"
		message += "Int"
	case "PurchaseItem":
		field = "purchaseItems"
		message = "Is required and must be array of items"
	case "ProductId":
		field = "productId"
		message += "Int"
	case "UnitPrice":
		field = "unitprice"
		message += "Float"
	case "Quantity":
		field = "quantity"
		message += "Float"
	case "PurchaseId":
		field = "purchaseId"
		message += "Int"

	}
	return field, message
}
