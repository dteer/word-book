package service

import (
	"fmt"
	"net"
	"word-book/config"
	util "word-book/service/pkg/utils"
	"word-book/service/recommend/strategy/one"
	service "word-book/service/service"

	"google.golang.org/grpc"
)

func Server() {
	// 刷新推荐
	var funcList = make([]func(), 1)
	funcList[0] = one.HandleRommend
	reFresh := util.TimingRefresh{
		F:     funcList,
		Times: config.C.Common.RemmandInterval,
	}
	go reFresh.Run()
	// 服务执行
	server := grpc.NewServer()
	service.RegisterProdServiceServer(server, service.ProdService)
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(fmt.Sprintf("服务器监听端口失败：%s", err))
	}
	fmt.Println("启动grpc服务器,监听端口: 8002")
	_ = server.Serve(listener)
}
