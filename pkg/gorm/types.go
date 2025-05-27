package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringArray []string

// 实现 driver.Valuer 接口，用于写入数据库
func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// 实现 sql.Scanner 接口，用于从数据库读取
func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSONArray")
	}
	return json.Unmarshal(bytes, s)
}
