/* единая логика кастомных ошибок */
package errors

import "fmt"

var (
	// rest api
	ErrorIsNotMatch  = New(0, "")
	StatusForbidden  = New(403, "Нет прав доступа")
	StatusBadRequest = New(400, "Некорректные параметры")
)

type Error struct {
	ErrorCode        int    `json:"code"`
	ErrorDescription string `json:"text"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("err_type:%d , err_des:%s", e.Code, e.ErrorDescription)
}
func (e *Error) Code() int {
	return e.ErrorCode
}

func New(code int, desc string) *Error {
	return &Error{
		ErrorCode:        code,
		ErrorDescription: desc,
	}
}
