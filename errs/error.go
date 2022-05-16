package errs

import "net/http"

func NotFoundError(msg string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}
func UnprocessableEntityError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	}
}

/* func ServerError(msg string) *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
} */

func UnAuthorizedError(msg string) *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}

type Error struct {
	Code    int            `json:",omitempty"`
	Message string         `json:"message"`
	Error   []ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

func ResponseError(msg string, code int) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func (e Error) AsResponse() *Error {
	return &Error{
		Code:    e.Code,
		Message: e.Message,
		Error:   e.Error,
	}
}

func (e Error) AsMessage() string {
	return e.Message
}
