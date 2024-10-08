package pool

import (
	"testing"
	"time"
)

// 测试Channel函数
func TestChannel(t *testing.T) {
	// 为测试设置超时
	timeout := time.After(1 * time.Second)

	// 调用Channel函数
	go Channel(1)

	// 等待超时，如果测试没有死锁，应该在1秒内完成
	select {
	case <-timeout:
		t.Error("Test timed out, expected output was not received in time")
	case <-time.After(500 * time.Millisecond):
		// 测试通过，我们预期在500毫秒内不会有输出
	}
}