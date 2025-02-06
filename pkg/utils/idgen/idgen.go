package idgen

import (
	"errors"
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sf     *sonyflake.Sonyflake
	sfOnce sync.Once
)

// // SFID 定义为 uint64 类型
// type SFID uint64

// // UUID 使用类型别名将 UUID 定义为 SFID
// type UUID = SFID

// // SFIDs 定义为 UUID 的切片
// type SFIDs []UUID

// GetSonyflake 返回一个单例的 Sonyflake 实例
func GetSonyflake() (*sonyflake.Sonyflake, error) {
	var err error
	sfOnce.Do(func() {
		settings := sonyflake.Settings{
			StartTime: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			MachineID: func() (uint16, error) {
				return 1, nil // 这里可以替换为获取机器ID的逻辑
			},
		}
		sf = sonyflake.NewSonyflake(settings)
		if sf == nil {
			err = errors.New("failed to create sonyflake instance")
		}
	})
	return sf, err
}

// GenerateUUID 生成一个 UUID
func GenerateUUID() (uint64, error) {
	sf, err := GetSonyflake()
	if err != nil {
		return 0, err
	}
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GenerateUUIDs 生成多个 UUIDs
func GenerateUUIDs(n int) ([]uint64, error) {
	var uuids []uint64
	for i := 0; i < n; i++ {
		uuid, err := GenerateUUID()
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, uuid)
	}
	return uuids, nil
}
