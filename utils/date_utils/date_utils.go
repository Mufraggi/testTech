package date_utils

import (
	"fmt"
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}


func GetNowString() time.Time {
	return FormatDateToTime(GetNow().Format(apiDateLayout))
}

func FormatDateToTime(date string) time.Time {
	res, err := time.Parse(apiDateLayout, date)
	if err != nil {
		fmt.Println(err)
	}
	return res
}