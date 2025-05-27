package conf

type MysqlConfig struct {
	// 示例 "user:pass@tcp(172.22.32.13:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DSN             string `json:"" yaml:"DSN"`
	MaxOpenConns    int    `json:"" yaml:"MaxOpenConns"`
	MaxIdleConns    int    `json:"" yaml:"MaxIdleConns"`
	ConnMaxLifetime int    `json:"" yaml:"ConnMaxLifetime"`
	MaxIdleTime     int    `json:"" yaml:"MaxIdleTime"`
}
