package main

import (
	"fmt"
	"time"
)

func lesson16() {
	time_now := time.Now()
	fmt.Println(time_now)
	fmt.Println(time_now.Year(), time_now.Month(), time_now.Day())

	time_utc := time_now.UTC()
	fmt.Println(time_utc)

	var one_week_nano time.Duration = 60 * 60 * 24 * 7 * 1e9
	one_week_later_from_now := time_now.Add(one_week_nano)
	fmt.Println(one_week_later_from_now)

	fmt.Println(time_now.Format(time.RFC822))
	fmt.Println(time_now.Format(time.ANSIC))
	//非常微妙的设计，固定使用这个日期+时间作为format的格式依据
	fmt.Println(time_now.Format("2006-01-02 15:04:05"))
}
