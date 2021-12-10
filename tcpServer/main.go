package main

import (
	"myEntry/pkg/log"
	"myEntry/tcpServer/common"
)

func main() {
	//服务初始化
	err := common.InitServer()
	if err != nil {
		log.Panic("Tcp Server Start Error : %s\n", err)
	}
}
