package main

import (
	"fmt"
	"reflect"
	"time"
)

type UserInfo struct {
	ID      int       `json:"id" redis:"id_redis"`
	Name    string    `json:"name" redis:"name_redis"`
	Created time.Time `json:"created" redis:"created_redis"`
}

func tt(t any) {
	types := reflect.TypeOf(t)
	fmt.Print(types.Elem())
	// // fmt.Print(types.Elem().)
	// fmt.Print(types.Elem().Field(0).Interface())
}

func main() {
	// run.InitLog()
	// config.InitConfig()
	// initiable.CheckDB()
	// initiable.InitGorm()
	// service.ServiceRun()

	// var user = &UserInfo{
	// 	ID:   1,
	// 	Name: "Jack",
	// }
	// util.GormFindJson(user)
	tt(1)
}
