package errcode

var (
	GameRequestErr = New(90001, "game request prams err")
	GameNoFoundErr = New(90002, "game no found")
)
