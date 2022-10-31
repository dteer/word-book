package main

import (
	"word-book/config"
	"word-book/run"
	"word-book/service"
)

func main() {
	// run.InitLog()
	config.InitConfig()
	run.InitGorm()
	service.Get()

}
