package main

import (
	"context"
	"log"
	"time"
	"word-book/config"
	"word-book/gui/service"

	"github.com/go-toast/toast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func shower(w *service.Data) {
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

func conn() *grpc.ClientConn {
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func main() {
	for {
		conn := conn()
		WordClient := service.NewProdServiceClient(conn)
		request := service.ShowRequest{
			Page:  1,
			Limit: 20,
		}
		words, err := WordClient.GetRecommendWord(context.Background(), &request)
		conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		for _, word := range words.Data {
			shower(word)
			time.Sleep(time.Duration(config.C.Common.Interval) * time.Second)
		}

	}

}
