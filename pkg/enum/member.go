package enum

// MemberType 会员类型
type MemberType int32

const (
	MemberDirect MemberType = iota
	MemberSub
)

// Index - 返回枚举项的字符值
func (m MemberType) Index() int32 {
	return int32(m)
}
func (m MemberType) GetType(index int32) MemberType {
	return MemberType(index)
}
