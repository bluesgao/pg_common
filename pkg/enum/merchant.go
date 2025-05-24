package enum

// MerchantStatus 自定义类型
type MerchantStatus int32

// 声明每个枚举项的索引值
const (
	MerchantStatusInit MerchantStatus = iota
	MerchantStatusEnable
	MerchantStatusDisable
)

// Index - 返回枚举项的字符值
func (w MerchantStatus) Index() int32 {
	return int32(w)
}

// MerchantType 商户类型
type MerchantType int32

const (
	MerchantSelf MerchantType = iota
	MerchantThird
)

// Index - 返回枚举项的字符值
func (m MerchantType) Index() int {
	return int(m)
}
func (m MerchantType) GetMerchantType(index int32) MerchantType {
	return MerchantType(index)
}
