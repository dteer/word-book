package util

import (
	"fmt"
	"io"
	"os"
	"time"

	"gorm.io/gorm"
)

/*
:func 获取今天0点的时间戳
:return 返回今天0点时间戳
*/
func GetNowDay() int {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	timeSamp := int(addTime.Unix())
	return timeSamp
}

/*
:func 获取相差的天数
:return 返回相差的天数
*/
func Interval(lastDay int, today int) int {
	startTime := time.Unix(int64(lastDay), 0)
	endTime := time.Unix(int64(today), 0)
	sub := int(endTime.Sub(startTime).Hours())
	days := sub / 24
	if (sub % 24) > 0 {
		days = days + 1
	}
	return days
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// todo 补充根据tag完成查询
func GormFindJson(a any, db *gorm.DB) {
}

func TimingRefresh(f func(), times int, ch <-chan struct{}) {
	waitTime := time.Duration(times) * time.Second
	timer := time.NewTimer(waitTime)
	for {
		select {
		case <-ch:
			return
		case <-timer.C:
			go f()
			timer.Reset(waitTime)
		}

	}
}
