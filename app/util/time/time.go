package time

import (
	"strconv"
	"time"
)

// 获取想要的当前时间格式
func GetCurrentTimeByFormat(format string) string {
	return time.Now().Format(format)
}

// 获取当前毫秒时间戳
func GetCurrentMilliTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// 将毫秒转化为想要的格式
func FormatUnixNanoTime(timestamp int64, format string) string {
	msString := strconv.FormatInt(timestamp, 10)
	msInt, err := strconv.ParseInt(msString, 10, 64)
	if err != nil {
		return time.Now().Format(format)
	}

	return time.Unix(0, msInt*int64(time.Millisecond)).Format(format)
}
