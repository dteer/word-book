package initiable

import (
	"log"
	"path"
	"sync"
	"time"
	"word-book/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormDB = make(map[string]*gorm.DB)

func ConnectGorm(name string) (*gorm.DB, error) {
	dns := path.Join(config.C.ItemPath, config.C.SQLite.DSN(name))
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		log.Fatal("连接mysql数据库失败, error=" + err.Error())
		return nil, err
	}
	sqlDB, _ := db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(50)           // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(300)          // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return db, nil
}

func InitGorm() {
	dbCofS := config.C.SQLite
	if len(dbCofS) == 0 {
		panic("无法连接数据库，请检查环境变量【ENV】和配置")
	}
	for dbName, _ := range dbCofS {
		db, err := ConnectGorm(dbName)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		GormDB[dbName] = db
	}
}

// GetGorm 获取GetGorm实例
func GetGorm(name string) *gorm.DB {
	var mu sync.RWMutex
	mu.Lock()
	db := GormDB[name]
	if db.Error != nil {
		db, _ = ConnectGorm(name)
		GormDB[name] = db
	}
	mu.Unlock()
	return db
}

// GetDefaultGorm 获取默认实例
func GetDefaultGorm() *gorm.DB {
	var mu sync.RWMutex
	mu.Lock()
	name := "default"
	db := GormDB[name]
	if db.Error != nil {
		db, _ = ConnectGorm(name)
		GormDB[name] = db
	}
	mu.Unlock()
	return db
}
