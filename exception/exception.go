package exception

import "fmt"

func NewNotFound(message string, a ...any) *ApiException {
	return New(404, message, a...)
}

func IsNotFound(err error) bool {

	if e, ok := err.(*ApiException); ok {
		if e.Code == 404 {
			return true
		}
	}
	return false
}

func NewAuthFailed(message string, a ...any) *ApiException {
	return New(500, message, a...)

}
func New(code int, message string, a ...any) *ApiException {
	return &ApiException{
		Code:    code,
		Message: fmt.Sprintf(message, a...),
	}
}

type ApiException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiException) Error() string {
	return e.Message
}
