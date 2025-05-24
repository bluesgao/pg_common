package id

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

type SfIdConfig struct {
	MachineID uint16 `json:",optional"`
}

func CreateSnowflakeIdGenerator(cfg SfIdConfig) *sonyflake.Sonyflake {
	settings := sonyflake.Settings{
		StartTime: time.Now(),
		MachineID: func() (uint16, error) {
			return cfg.MachineID, nil
		},
	}
	sf := sonyflake.NewSonyflake(settings)
	if sf == nil {
		panic("sony flake not created")
	}
	return sf
}

// GenerateFixedLengthIDWithSnowflake 使用雪花算法生成固定长度 ID
func GenerateFixedLengthIDWithSnowflake(length int) (string, error) {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		return "", fmt.Errorf("failed to initialize Sonyflake")
	}

	// 生成雪花 ID
	id, err := sf.NextID()
	if err != nil {
		return "", fmt.Errorf("failed to generate ID: %v", err)
	}

	// 格式化为固定长度
	idStr := fmt.Sprintf("%d", id)
	if len(idStr) > length {
		idStr = idStr[:length]
	} else if len(idStr) < length {
		idStr = fmt.Sprintf("%0*s", length, idStr)
	}

	return idStr, nil
}
