package errs


import (
"fmt"
)

// ErrorType 错误码类型 包括框架错误码和业务错误码
const (
	ErrorTypeFramework = "framework"
	ErrorTypeBusiness  = "business"
)


const (
	ErrorOK = 0
	ErrorUnKnow = 1000
)

type Error struct {
	Type string `json:"type,omitempty"`
	Code int32 `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (e *Error) Error() string {
	if e == nil {
		return "nil"
	}
	if e.Type == ErrorTypeFramework {
		return fmt.Sprintf("type:framework, code:%d, msg:%s", e.Code, e.Msg)
	}
	return fmt.Sprintf("type:business, code:%d, msg:%s", e.Code, e.Msg)
}

// New 创建一个error，默认为业务错误类型，提高业务开发效率
func New(code int, msg string) *Error {
	return &Error{
		Type: ErrorTypeBusiness,
		Code: int32(code),
		Msg:  msg,
	}
}

// NewFrameError 创建一个框架error
func NewFrameError(code int, msg string) *Error {
	return &Error{
		Type: ErrorTypeFramework,
		Code: int32(code),
		Msg:  msg,
	}
}

func NewSuccess() *Error {
	return &Error{
		Code: 0,
		Msg:  "success",
	}
}

