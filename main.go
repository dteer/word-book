package main

import (
	"word-book/config"
	"word-book/dao/word"
	"word-book/initiable"
	"word-book/service"
)

func main() {
	config.InitConfig()
	initiable.CheckDB()
	initiable.InitGorm()
	word.UdateWordComplexity()
	service.ServiceRun()
}
