package dto

import (
	"net/http"

	"github.com/brackbk/purchase-service/errs"
	"github.com/go-playground/validator/v10"
)

type baseDto interface {
	ErrorMessage(s string) (string, string)
}

func Validate(b baseDto) *errs.Error {
	v := validator.New()
	err := v.Struct(b)
	var errList []errs.ErrorMessage

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			field, message := b.ErrorMessage(e.StructField())
			errList = append(errList, errs.ErrorMessage{
				Field:   field,
				Message: message,
			})
		}
		return &errs.Error{
			Code:    http.StatusBadRequest,
			Message: "This Fields below must be filled",
			Error:   errList,
		}
	}
	return nil
}
