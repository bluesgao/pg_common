package mq

// topic 命名规范
// 1. 以 pg_ 开头，表示pg系统的消息
// 2. pg_XXX_YYY 表示pg系统中 XXX 模块的 YYY 消息

var (
	TOPIC_USER_REGISTER     string = "pg_user_register"
	TOPIC_MERCHANT_REGISTER string = "pg_merchant_register"
	TOPIC_AGENT_REGISTER    string = "pg_agent_register"
)
