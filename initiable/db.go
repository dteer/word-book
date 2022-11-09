package initiable

import (
	"os"
	"path"
	"word-book/config"
	util "word-book/pkg/utils"
)

// 处理数据库相关的

// 检查db文件是否存在，否则获取init的
func CheckDB() {
	dbFile := path.Join(config.C.ItemPath, config.C.SQLite.Default().File)
	initDBFile := path.Join(config.C.ItemPath, config.C.SQLite.Default().InitFile)
	if _, err := os.Stat(dbFile); err != nil {
		if os.IsNotExist(err) {
			// 复制初始文件
			util.CopyFile(initDBFile, dbFile)
		}
	} else {
		recover()
	}
}
