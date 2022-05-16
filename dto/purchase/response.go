package purchase

import (
	"github.com/brackbk/purchase-service/domain"
	"github.com/brackbk/purchase-service/dto/purchaseItem"
)

type ResponseDTO struct {
	Id            uint                       `jsaon:"id"`
	PaymentStatus string                     `json:"paymentStatus"`
	PaymentType   string                     `json:"paymentType"`
	InvoiceNumber string                     `json:"invoiceNumber"`
	CompanyId     *int                       `json:"companyId"`
	ProviderId    *int                       `json:"providerId"`
	LocationId    *int                       `json:"locationId"`
	CurrencyId    *int                       `json:"currencyId"`
	Total         *float64                   `json:"total"`
	PurchaseItem  []purchaseItem.ResponseDTO `json:"purchaseitems"`
}

func MountResponse(domain domain.Purchase) ResponseDTO {
	var purchaseItems []purchaseItem.ResponseDTO
	for _, item := range domain.PurchaseItem {
		itemResponse := purchaseItem.ResponseDTO{
			Id:         item.ID,
			PurchaseId: item.PurchaseId,
			ProductId:  item.ProductId,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.TotalPrice,
			CurrencyId: item.CurrencyId,
		}
		purchaseItems = append(purchaseItems, itemResponse)
	}
	return ResponseDTO{
		Id:            domain.ID,
		PaymentStatus: domain.PaymentStatus,
		PaymentType:   domain.PaymentType,
		InvoiceNumber: domain.InvoiceNumber,
		CompanyId:     domain.CompanyId,
		ProviderId:    domain.ProviderId,
		CurrencyId:    domain.CurrencyId,
		Total:         domain.Total,
		PurchaseItem:  purchaseItems,
	}
}
