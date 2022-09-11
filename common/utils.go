package common

import (
	"strconv"
	"time"
)

func ConvertTimeStampToTime(timeStamp string) time.Time {
	i, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.UnixMilli(i)
}
