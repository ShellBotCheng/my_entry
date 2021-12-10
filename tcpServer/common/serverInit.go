package common

import (
	"github.com/spf13/pflag"
	"myEntry/pkg/log"
	"myEntry/pkg/mysql"
	"myEntry/pkg/redis"
	"myEntry/pkg/rpc"
	"myEntry/tcpServer/conf"
)

var (
	tcpCfgFile = pflag.StringP("config", "c", "./tcpServer/conf/dev.yml", "config file path.")
)

func InitServer() error {
	err := log.Init()
	if err != nil {
		return err
	}
	// init config
	cfg := conf.Init(*tcpCfgFile)

	// init redis
	redis.Init(&cfg.Redis)

	// init mysql
	mysql.Init(&cfg.Mysql)

	// init rpcServer
	addr := cfg.Tcp.Addr
	server := rpc.InitRpcServer(addr)
	log.Info("Start Tcp Server Success By Addr: [ %s ] ", addr)
	go server.Run()
	select {}
}
