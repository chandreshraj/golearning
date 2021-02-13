package runtimeUtilities

import "time"

func Microseconds() int64 {
	return time.Now().Local().Unix()
}
