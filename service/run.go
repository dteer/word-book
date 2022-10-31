package service

import (
	"fmt"
	"word-book/dao/word"
)

func Get() {
	data := word.Find()
	fmt.Println(data)
}
