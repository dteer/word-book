package service

import (
	"word-book/dao/word"
)

func handleAddWord(count int) (words []word.Word) {
	StartTime := 0
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		StartTime: &StartTime,
		Order:     "id asc",
	}
	words = word.Find(findData)
	return
}

func getNewWord(count int, tody int) (words []word.Word) {
	// 获取今天需要学习的数据
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		StartTime: &tody,
	}
	words = word.Find(findData)
	// 如果今天还没学习，需要录入信息
	if len(words) < count {
		addWords := handleAddWord(count - len(words))
		word.SetTodayStartTime(addWords, tody)
		words = append(words, addWords...)
	}
	return
}
