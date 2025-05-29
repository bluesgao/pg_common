package errcode

var (
	DomainIsInvalid         = New(990001, "domain is invalid")
	InviteCodeIsInvalid     = New(990002, "invite code is invalid")
	ReqParamsErr            = New(990003, "request params err")
	UserIsInvalid           = New(990004, "user is invalid")
	UserTokenIsInvalid      = New(990005, "user token is invalid")
	UserTokenNotFound       = New(990006, "user token not found")
	UserTokenIllegal        = New(990007, "user token is illegal")
	UserRefreshTokenIllegal = New(990008, "user refresh token is illegal")
	RpcCallFailed           = New(990009, "rpc call failed")
	UserAuthErr             = New(990010, "user auth failed")
)
