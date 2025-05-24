package enum

type AccountStatus int32

// 声明每个枚举项的索引值
const (
	AccountStatusInit MerchantStatus = iota
	AccountStatusEnable
	AccountStatusDisable
)

// Index - 返回枚举项的字符值
func (w AccountStatus) Index() int32 {
	return int32(w)
}

type AccountTransType int32

const (
	AccountTransTypeDeposit AccountTransType = iota + 1
	AccountTransTypeWithdraw
	AccountTransTypeTransfer
	AccountTransTypeFreeze
	AccountTransTypeUnfreeze
	AccountTransTypeAdjustment
)

// Index - 返回枚举项的字符值
func (w AccountTransType) Index() int32 {
	return int32(w)
}

// OwnerType 钱包所有者类型
type OwnerType int32

const (
	OwnerMerchant OwnerType = iota + 1
	OwnerAgent
	OwnerMember
)

// Index - 返回枚举项的字符值
func (m OwnerType) Index() int32 {
	return int32(m)
}

// DcType 借贷标识
type DcType int32

const (
	//借记（Debit）
	Debit DcType = iota + 1
	//贷记（Credit）
	Credit
)

// Index - 返回枚举项的字符值
func (m DcType) Index() int32 {
	return int32(m)
}
