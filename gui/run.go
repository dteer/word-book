package gui

import (
	"log"
	"word-book/service/dao/word"

	"github.com/go-toast/toast"
)

func showWord(w word.Word) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   w.Title,
		Message: "[" + w.PhoneticStymbol + "]\n\n" + w.Description,
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		// Actions: []toast.Action{
		// 	{"protocol", "按钮1", "https://www.google.com/"},
		// 	{"protocol", "按钮2", "https://github.com/"},
		// },
		Duration: "long",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
