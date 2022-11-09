package config

import (
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var C Config

func LoadConfig(conf *Config, confPath string, filteType string) {
	var confObj = viper.New()
	confObj.SetConfigType(filteType)
	confObj.AddConfigPath(confPath)
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
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(path.Dir(filename))
	}
	var confPath string
	if env == "" {
		confPath = path.Join(abPath, "config/dev")
	} else if env == "PROD" {
		confPath = path.Join(abPath, "config/prod")
	} else {
		panic("请设置环境变量: ENV=DEBUG(测试) ENV=PROD(正式)")
	}
	LoadConfig(&C, confPath, "yaml")
	C.ItemPath = abPath
}

type Config struct {
	ItemPath    string
	RunMode     string     // 启动环境 debug/prod
	PrintConfig bool       // 配置打印
	Redis       Redis      // redis配置
	SQLite      SQLiteList // mysql配置
	Common      Common     // 常规操作
}

func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

type SQLiteList map[string]struct {
	File     string
	InitFile string
}

func (a SQLiteList) DSN(name string) string {
	return a[name].File
}
func (a SQLiteList) Default() struct {
	File     string
	InitFile string
} {
	return a["default"]
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
	Interval        int
	RemmandInterval int
	New             int
	Old             int
}
