package mq

type UserRegisterMsg struct {
	MerchantId   string `json:"merchant_id"`
	AgentId      string `json:"agent_id"`
	UserId       string `json:"user_id"`
	RegisterType int32  `json:"register_type"`
	RegisterNo   string `json:"register_no"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	Gender       string `json:"gender"`
	Birthday     string `json:"birthday"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	Address      string `json:"address"`
	UserType     int32  `json:"user_type"`
	Avatar       string `json:"avatar"`
	InviteCode   string `json:"invite_code"`
	Password     string `json:"password"`
}
