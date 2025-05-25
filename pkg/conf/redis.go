package conf

type RedisConfig struct {
	Password string `json:",optional"`
	//xx.xx.xx.xx:6379
	Host string `json:",optional"`

	//MaxIdle     int   `json:",optional"`
	//MaxActive   int   `json:",optional"`
	//IdleTimeout int64 `json:",optional"`
}
