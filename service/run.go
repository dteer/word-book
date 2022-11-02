package service

import (
	"time"
	"word-book/config"
	"word-book/dao/word"
	util "word-book/pkg/utils"
)

func ServiceRun() {
	newCount := config.C.Common.New
	oldCount := config.C.Common.Old
	Today := util.GetNowDay()
	newWords := getNewWord(newCount, Today)
	recommendWords := getRecommendWrods(oldCount)
	words := append(newWords, recommendWords...)
	for {
		for index := range words {
			w := words[index]
			popWord(w)
			word.UdateShowTime(w)
			words[index].ShowTime += 1
			time.Sleep(time.Duration(config.C.Common.Interval) * time.Second)
		}
	}
}
