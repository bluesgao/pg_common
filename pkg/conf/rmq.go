package conf

type RocketMqConfig struct {
	NameServer    []string `json:"nameServer"`
	ProducerGroup string   `json:"producerGroup"`
	ConsumerGroup string   `json:"consumerGroup"`
}
