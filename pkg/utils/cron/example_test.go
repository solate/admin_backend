package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCronManager(t *testing.T) {
	// 创建cron管理器
	cm := NewCronManager()
	defer cm.Stop()

	// 添加定时任务
	jobID := "test_job"
	_, err := cm.AddJob(jobID, "*/5 * * * * *", func() error {
		fmt.Println("执行定时任务:", time.Now())
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	// 等待任务执行
	time.Sleep(10 * time.Second)

	// 移除任务
	if removed := cm.RemoveJob(jobID); !removed {
		t.Fatal("移除任务失败")
	}
}
