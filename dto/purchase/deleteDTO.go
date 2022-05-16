package purchase

type DeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

func (d DeleteRequest) ErrorMessage(s string) (string, string) {
	var field string
	message := "is required and must be of type"
	switch s {
	case "Id":
		field = "id"
		message += "Int"
	}
	return field, message
}
