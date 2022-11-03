package util

import (
	"fmt"
	"io"
	"os"
	"reflect"
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

func GormFindJson(a any, db *gorm.DB) {
	types := reflect.TypeOf(a)
	values := reflect.ValueOf(a)
	structLen := types.Elem().NumField()
	for i := 0; i < structLen; i++ {
		fields := types.Elem().Field(i)
		value := values.Elem().Field(i).Interface()

		filed := fields.Tag.Get("filed")
		abc := fields.Tag.Get("abc")
		condition := fields.Tag.Get("condition")
		// defaultValue := fields.Tag.Get("default")
		foctor := fmt.Sprintf("%s %s ?", filed, condition)
		if abc == "nil" {
			if value != nil {
				db = db.Where(foctor, value)
			}
		} else {
			filedType := reflect.TypeOf(fields)
			filedValue := reflect.ValueOf(values.Interface())
			switch filedValue.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
				value = int(filedValue)
			case reflect.String:

			}
			if string(value) != abc {
				db = db.Where(foctor, value)
			}
		}
	}
	field1 := types.Elem().Field(0)
	field2 := types.Elem().Field(1)
	tagName1 := field1.Tag.Get("json")  //  获得字段 ID 的json值
	tagName2 := field2.Tag.Get("redis") //  获得字段 Name 的redis值
	fmt.Println(tagName1, tagName2)
	fmt.Print(values.Elem().Field(0).Interface()) // 获取值
}
