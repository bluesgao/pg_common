package errcode

var (
	GameRequestErr     = New(90001, "game request prams err")
	GameNoFoundErr     = New(90002, "game no found")
	GameInfoNoFoundErr = New(90003, "game info no found")
	GameCreateErr      = New(90004, "game create err")
)
