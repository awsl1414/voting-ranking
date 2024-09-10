/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-19 23:20:44
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-09-10 16:58:22
 * @FilePath: /voting-ranking/common/util/times.go
 * @Description:
 *
 */
package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// HTime 自定义时间类型，继承自 time.Time，用于处理时间格式的序列化与反序列化
type HTime struct {
	time.Time
}

var (
	// 定义时间格式
	formatTime = "2006-01-02 15:04:05"
)

// MarshalJSON 实现 JSON 序列化，将时间格式化为指定的字符串形式
func (t HTime) MarshalJSON() ([]byte, error) {
	// 将时间格式化为 "2006-01-02 15:04:05" 的字符串，并加上双引号
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现 JSON 反序列化，将字符串解析为时间
func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	*t = HTime{Time: now}
	return
}

// Value 实现数据库存储接口，处理时间的存储
func (t HTime) Value() (driver.Value, error) {
	// 判断是否为零时间，如果是则返回 nil
	var zeroTime time.Time

	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现数据库读取接口，将数据库中的时间值赋给 HTime 类型
func (t *HTime) Scan(v interface{}) error {
	// 检查数据库值是否是 time.Time 类型
	value, ok := v.(time.Time)
	if ok {
		// 如果是，则将其赋值给 HTime 类型
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
