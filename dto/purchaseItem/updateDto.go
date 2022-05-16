package purchaseItem

type UpdateRequest struct {
	Id         int     `json:"id" validate:"required,gt=0"`
	PurchaseId int     `json:"purchaseId" validate:"required,gt=0"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unitprice"`
}

func (u UpdateRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "Is required and must be of type: "
	switch s {
	case "Id":
		field = "id"
		message += "Int"
	case "PurchaseId":
		field = "purchaseId"
		message += "Int"
	case "Quantity":
		field = "quantity"
		message = "Is required and must be greater than 0"
	case "UnitPrice":
		field = "unitprice"
		message += "float"
	}
	return field, message
}
