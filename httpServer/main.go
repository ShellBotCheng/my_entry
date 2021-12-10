package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"myEntry/httpServer/common"
	"myEntry/httpServer/conf"
	"myEntry/httpServer/routers"
	"myEntry/pkg/log"
	"net/http"
)

var (
	httpCfgFile = pflag.StringP("config", "c", "./httpServer/conf/dev.yml", "config file path.")
)

func main() {
	err := log.Init()
	if err != nil {
		fmt.Printf("Http Server Start Error : %s\n", err)
	}

	pflag.Parse()

	// init config
	cfg := conf.Init(*httpCfgFile)

	// init tcpPool
	err = common.TcpPoolInit(&cfg.TcpPool)
	if err != nil {
		log.Panic("Http Server Start Error : %s\n", err)
	}

	addr := cfg.Http.Addr
	srv := &http.Server{
		Addr:    addr,
		Handler: routers.InitRouter(),
	}
	log.Info("start http server addr: [ %s ] ", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("server error: %s", err)
	}
}
