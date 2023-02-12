package services

import "time"

func SetNowDate() int64 {
	return time.Now().Unix()
}

func SetTokenExpireAt(duration int64) int64 {
	return time.Now().Add(time.Duration(duration)).Unix()
}
