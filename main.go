package main

import (
	"flag"
	"go-shop/config"
	"go-shop/dao/mysql"
	"go-shop/handler"
	"go-shop/logger"
	"go-shop/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// 加载配置文件
	var cfn string
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件")
	err := config.Init(cfn)
	if err != nil {
		panic(err)
	}

	// 加载日志
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		panic(err)
	}

	// 初始化mysql
	err = mysql.Init(config.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}

	// 初始化Consul

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 创建grpc服务
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGoodsServer(s, &handler.GoodsSrv{})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
