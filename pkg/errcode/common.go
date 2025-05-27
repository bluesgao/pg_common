package errcode

// common code
var (
	SUCCESS = New(0, "success")
)

// errorcode组成规则(6位长度)
// 1. 1-9 为系统错误码
// 2. 01-99 为业务错误码
//01-merchant
//02-agent
//03-user
//04-order
//05-wallet
//06-bill
//07-pay
//08-paygateway
//09-game
//99-apigateway

// 3. 001-999 为业务错误码
