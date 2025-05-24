package errcode

var (
	AgentIdIsNull                 = New(202001, "agent id is null")
	AgentNameIsNull               = New(202002, "agent name is null")
	AgentPhoneIsNull              = New(202003, "agent phone is null")
	AgentEmailIsNull              = New(202004, "agent email is null")
	AgentAvatarIsNull             = New(202005, "agent avatar is null")
	AgentCreateErr                = New(202006, "agent create err")
	AgentExistsErr                = New(202007, "agent already exists")
	AgentNoFoundErr               = New(202008, "agent no found")
	AgentUpdateReqParamsErr       = New(202009, "agent update req params err")
	AgentUpdateErr                = New(202010, "agent update err")
	AgentCreateReqParamsErr       = New(202011, "agent create req params err")
	AgentUpdateStatusReqParamsErr = New(202012, "agent update status req params err")
	AgentUpdateStatusErr          = New(202013, "agent update status err")
	AgentStatusNotActiveErr       = New(202014, "agent status not active")
)
