package config

import (
	"os"

	"github.com/spf13/viper"
)

var C Config

func LoadConfig(conf *Config, path string, filteType string) {
	var confObj = viper.New()
	confObj.SetConfigType(filteType)
	confObj.AddConfigPath(path)
	//读取配置文件内容
	if err := confObj.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := confObj.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

func InitConfig() {
	env := os.Getenv("ENV")
	var path string
	if env == "" {
		path = "config/dev"
	} else if env == "PROD" {
		path = "config/prod"
	} else {
		panic("请设置环境变量: ENV=DEBUG(测试) ENV=PROD(正式)")
	}
	LoadConfig(&C, path, "yaml")
}

type Config struct {
	RunMode     string // 启动环境 debug/prod
	PrintConfig bool   // 配置打印
	Redis       Redis  // redis配置
	SQLite      SQLite // mysql配置
	Common      Common // 常规操作
}

func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

type SQLite map[string]struct {
	File string
}

func (a SQLite) DSN(name string) string {
	return a[name].File
}

type Redis map[string]struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       int
	Timeout  int
}

type Common struct {
	Interval int
}
