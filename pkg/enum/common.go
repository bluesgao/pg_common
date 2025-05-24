package enum

const (
	MaxRetries = 3
)

const (
	// MerchantIdPrefix 商户
	MerchantIdPrefix = 1
	// AgentIdPrefix 代理
	AgentIdPrefix = 2
	// MemberIdPrefix 会员
	MemberIdPrefix = 3
	// WalletIdPrefix 钱包
	WalletIdPrefix = 8
)

// SystemType 系统标识
type SystemType int32

const (
	SystemOrder SystemType = iota + 10
	SystemPay
)

// Index - 返回枚举项的字符值
func (m SystemType) Index() int32 {
	return int32(m)
}

// BizType 业务标识
type BizType int32

const (
	//游戏订单
	BizGameOrder BizType = iota + 10
	//支付订单
	BizPayOrder
)

// Index - 返回枚举项的字符值
func (m BizType) Index() int32 {
	return int32(m)
}
