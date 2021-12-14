package common

import (
	"github.com/spf13/pflag"
	"myEntry/pkg/log"
	"myEntry/pkg/lru"
	"myEntry/pkg/mysql"
	"myEntry/pkg/redis"
	"myEntry/pkg/rpc"
	"myEntry/tcpServer/conf"
)

var (
	tcpCfgFile = pflag.StringP("config", "c", "./tcpServer/conf/dev.yml", "config file path.")
)

// Cfg global config
var Cfg *conf.Config

func InitServer() error {
	err := log.Init()
	if err != nil {
		return err
	}
	// init config
	Cfg = conf.Init(*tcpCfgFile)
	// init redis
	redis.Init(&Cfg.Redis)

	// init mysql
	mysql.Init(&Cfg.Mysql)

	// init lru cache
	lru.InitLru(Cfg.Lru.Max)

	// init rpcServer
	addr := Cfg.Tcp.Addr
	server := rpc.InitRpcServer(addr)
	log.Info("Start Tcp Server Success By Addr: [ %s ] ", addr)
	go server.Run()
	select {}
}
