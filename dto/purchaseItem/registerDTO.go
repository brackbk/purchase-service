package purchaseItem

type RegisterRequest struct {
	Id         uint    `json:"id"`
	PurchaseId int     `json:"purchaseId" validate:"omitempty,gt=0"`
	ProductId  int     `json:"productId" validate:"gt=0"`
	Quantity   int     `json:"quantity" validate:"gt=0"`
	UnitPrice  float64 `json:"unitprice" validate:"gt=0"`
}

func (r RegisterRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "Is required and must be of type:"
	switch s {
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

type ResponseDTO struct {
	Id         uint    `jsaon:"id"`
	PurchaseId int     `json:"purchaseid"`
	ProductId  int     `json:"productid"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unitprice"`
	TotalPrice float64 `json:"totalprice"`
	CurrencyId int     `json:"currencyid"`
}
