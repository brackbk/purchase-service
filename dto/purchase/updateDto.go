package purchase

import (
	"time"

	"github.com/brackbk/purchase-service/errs"
)

type UpdateRequest struct {
	Id            int                       `json:"id" validate:"required"`
	PaymentType   string                    `json:"paymentType"`
	ProviderId    *int                      `json:"providerId" validate:"omitempty,gt=0"`
	LocationId    *int                      `json:"locationId" validate:"omitempty,gt=0"`
	PurchaseDate  *time.Time                `json:"purchaseDate"`
	PaymentStatus string                    `json:"paymentStatus"`
	PurchaseItem  []*UpdateWithItemsRequest `json:"purchaseItems" validate:"omitempty,min=1,dive,required"`
	InvoiceNumber string                    `json:"invoiceNumber"`
}

type UpdateWithItemsRequest struct {
	Id        int     `json:"id"`
	ProductId int     `json:"productId" validate:"required,gt=0"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	UnitPrice float64 `json:"unitprice" validate:"required,gt=0"`
}

func (u UpdateRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "Is required and must be of type:"
	switch s {
	case "Id":
		field = "id"
		message += "Int"
	case "ProviderId":
		field = "providerId"
		message += "Int"
	case "LocationId":
		field = "locationId"
		message += "Int"
	case "PurchaseItem":
		field = "purchaseItems"
		message += "array of Items"
	case "UnitPrice":
		field = "unitprice"
		message += "Float"
	case "Quantity":
		field = "quantity"
		message += "Float"
	case "PurchaseId":
		field = "purchaseId"
		message += "Int"
	case "ProductId":
		field = "productId"
		message += "Int"
	}

	return field, message
}
func UpdatePurchase(
	Id int,
	PaymentStatus string,
	PaymentType string,
	ProviderId *int,
	LocationId *int,
) (*UpdateRequest, *errs.Error) {

	purchase := &UpdateRequest{
		Id:            Id,
		PaymentStatus: PaymentStatus,
		PaymentType:   PaymentType,
		ProviderId:    ProviderId,
		LocationId:    LocationId,
	}
	return purchase, nil
}
