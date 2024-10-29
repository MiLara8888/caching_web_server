/* единая логика кастомных ошибок */
package errors

import "fmt"

var (
	// rest api
	ErrorIsNotMatch         = New(0, "")

)


type Error struct {
	errorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("err_type:%d , err_des:%s", e.Code, e.ErrorDescription)
}
func (e *Error) Code() int {
	return e.errorCode
}

func New(code int, desc string) *Error {
	return &Error{
		errorCode:        code,
		ErrorDescription: desc,
	}
}
