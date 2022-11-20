package main

import (
	"word-book/initiable"
	"word-book/service"
	"word-book/service/dao/word"
)

func main() {
	initiable.InitGorm()
	word.UdateWordComplexity()
	service.ServiceRun()
}
