package time

import (
	"testing"
	"time"
)

func TestNowAddDay(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 增加5天
	dayToAdd := 5
	expected := now.AddDate(0, 0, dayToAdd)
	actual := NowAddDay(dayToAdd)

	// 检查函数返回的时间是否与预期相等
	if !actual.Truncate(time.Second).Equal(expected.Truncate(time.Second)) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 增加0天的测试
	dayToAdd = 0
	expected = now.AddDate(0, 0, dayToAdd)
	actual = NowAddDay(dayToAdd)

	// 检查函数返回的时间是否与预期相等
	if !actual.Truncate(time.Second).Equal(expected.Truncate(time.Second)) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 增加负数天数的测试
	dayToAdd = -10
	expected = now.AddDate(0, 0, dayToAdd)
	actual = NowAddDay(dayToAdd)

	// 检查函数返回的时间是否与预期相等
	if !actual.Truncate(time.Second).Equal(expected.Truncate(time.Second)) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
