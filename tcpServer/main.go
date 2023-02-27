package main

import (
	"fmt"
	"hash/crc32"
	"myEntry/pkg/log"
	"myEntry/tcpServer/common"
	"strconv"
)

func main() {
	//服务初始化
	err := common.InitServer()
	if err != nil {
		log.Panic("Tcp Server Start Error : %s\n", err)
	}
	table := getTableName("admin")
	fmt.Printf("table : %s", table)
}

func getTableName(uname string) string {
	return fmt.Sprintf("user_%04s", strconv.Itoa(int(crc32.ChecksumIEEE([]byte(uname))%128)))
}
