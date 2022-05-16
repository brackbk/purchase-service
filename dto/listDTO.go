package dto

type ListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit" validate:"gte=1,lte=20"`
}

func (t ListRequest) ErrorMessage(s string) (string, string) {
	var field string
	var message string
	switch s {
	case "Limit":
		field = "limit"
		message = "Is required and must be upper than 0 and under or equal to 20"
	}
	return field, message
}
