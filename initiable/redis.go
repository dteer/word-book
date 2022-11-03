package initiable

import (
	"fmt"
	"log"
	"time"
	"word-book/config"

	redigo "github.com/gomodule/redigo/redis"
)

var redisDB = make(map[string]*redigo.Pool)

func ConnectRedis(name string) (*redigo.Pool, error) {
	redisConfS := config.C.Redis
	redisConf := redisConfS[name]
	addr := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	redis := &redigo.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		MaxActive:   1000, //最大连接数
		Wait:        true, //接池链接数达到上限时，会阻塞等待其他协程用完归还之后继续执行
		Dial: func() (redigo.Conn, error) {
			redisConn, err := redigo.Dial("tcp", addr)
			if err != nil {
				panic(err)
				return nil, err
			}
			if redisConf.Password != "" {
				if _, err := redisConn.Do("AUTH", redisConf.Password); err != nil {
					redisConn.Close()
				}
			}
			if _, err := redisConn.Do("SELECT", redisConf.Db); err != nil {
				redisConn.Close()
				panic(err)
			}
			_, err = redisConn.Do("PING")
			return redisConn, nil
		},
	}
	return redis, nil
}

func InitRedis() map[string]*redigo.Pool {
	var redisDB = make(map[string]*redigo.Pool)
	redisConfS := config.C.Redis
	if len(redisConfS) == 0 {
		panic("无法连接redis，请检查环境变量【ENV】和配置")
	}
	for redisName, _ := range redisConfS {
		redis, err := ConnectRedis(redisName)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		redisDB[redisName] = redis
	}
	return redisDB
}

// GetRedis 获取redis实例
func GetRedis(name string) *redigo.Pool {
	redis, ok := redisDB[name]
	if !ok {
		return nil
	}
	return redis
}
