package service

import (
	"log"
	"time"
	"word-book/config"
	"word-book/dao/word"

	"github.com/go-toast/toast"
)

func getNewWord(count int, tody int) (words []word.Word) {
	// 获取今天需要学习的数据
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		StartTime: tody,
		ShowTime:  -1,
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

func handleAddWord(count int) (words []word.Word) {
	findData := word.FindData{
		Page:      1,
		Limit:     count,
		ShowTime:  0,
		StartTime: 0,
	}
	words = word.Find(findData)
	return
}

func getOldWord(count int, tody int) (words []word.Word) {
	findData := word.FindData{
		Page:        1,
		Limit:       count,
		StartTimeLt: tody,
		Order:       "start_time desc",
	}
	words = word.Find(findData)
	return
}

func popWord(w word.Word) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   w.Title + "  [" + w.PhoneticStymbol + "]",
		Message: w.Description,
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "按钮1", "https://www.google.com/"},
			{"protocol", "按钮2", "https://github.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func Run() {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	timeSamp := int(addTime.Unix())
	newWords := getNewWord(10, timeSamp)

	oldWords := getOldWord(20, timeSamp)

	words := append(newWords, oldWords...)
	for {
		for _, w := range words {
			popWord(w)
			word.UdateShowTime(w)
			time.Sleep(time.Duration(config.C.Common.Interval) * time.Second)
		}
	}

}
