package purchaseItem

type DeleteRequest struct {
	Id         int `json:"id" validate:"required"`
	PurchaseId int `json:"purchaseId" validate:"gt=0"`
}

func (d DeleteRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "Is required and must be of type: "
	switch s {
	case "Id":
		field = "id"
		message += "Int"
	case "PurchaseId":
		field = "purchaseId"
		message += "Int"
	}
	return field, message
}
