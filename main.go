package main

import (
	"os"
	"word-book/config"
	"word-book/initiable"
	util "word-book/pkg/utils"
	"word-book/service"
)

// 检查db文件是否存在，否则获取init的
func checkDB() {
	dbFile := config.C.SQLite.Default().File
	initDBFile := config.C.SQLite.Default().InitFile
	if _, err := os.Stat(dbFile); err != nil {
		if os.IsNotExist(err) {
			// 复制初始文件
			util.CopyFile(initDBFile, dbFile)
		}
	} else {
		recover()
	}
}

func main() {
	// run.InitLog()
	config.InitConfig()
	checkDB()
	initiable.InitGorm()
	service.ServiceRun()
}
