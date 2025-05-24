package enum

// AgentType 代理类型
type AgentType int32

const (
	AgentDirect AgentType = iota
	AgentSub
)

// Index - 返回枚举项的字符值
func (m AgentType) Index() int32 {
	return int32(m)
}
func (m AgentType) GetType(index int32) AgentType {
	return AgentType(index)
}

// AgentStatus 自定义类型
type AgentStatus int32

// 声明每个枚举项的索引值
const (
	AgentStatusInit AgentStatus = iota
	AgentStatusEnable
	AgentStatusDisable
)

// Index - 返回枚举项的字符值
func (w AgentStatus) Index() int32 {
	return int32(w)
}
