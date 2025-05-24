package errcode

import "fmt"

type ErrCode struct {
	code int32
	msg  string
}

func (e *ErrCode) Error() string {
	return fmt.Sprintf("%d-%s", e.code, e.msg)
}

func (e *ErrCode) Code() int32 {
	return e.code
}
func (e *ErrCode) Msg() string {
	return e.msg
}

func New(code int32, msg string) *ErrCode {
	return &ErrCode{code, msg}
}
