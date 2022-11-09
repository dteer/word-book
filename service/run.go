package service

import (
	"time"
	"word-book/config"
	"word-book/dao/word"
	util "word-book/pkg/utils"
	"word-book/recommend/strategy/one"
)

func ServiceRun() {
	// 刷新推荐
	go util.TimingRefresh(one.HandleRommend, config.C.Common.RemmandInterval, make(<-chan struct{}))

	newCount := config.C.Common.New
	oldCount := config.C.Common.Old
	Today := util.GetNowDay()
	newWords := getNewWord(newCount, Today)
	wordChanel := make(chan word.Word, 20)
	go handle(wordChanel)
	for {
		words := append(newWords, one.GetRecommendWrods(oldCount)...)
		for _, w := range words {
			popWord(w)
			wordChanel <- w
			time.Sleep(time.Duration(config.C.Common.Interval) * time.Second)
		}
	}
}

// 通过队列进行处理相关信息
func handle(wordLi chan word.Word) {
	data := <-wordLi
	word.UdateShowTime(data)
}
