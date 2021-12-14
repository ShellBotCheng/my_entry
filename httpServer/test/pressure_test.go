package test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

var (
	no      = 0
	ok      = 0   //记录请求成功失败数
	useTime = 0.0 //使用时间
	num     = 100 //并发个数
)

// BenchmarkLogin 登录压测
func BenchmarkLogin(t *testing.B) {
	startTime := time.Now().UnixNano() //记录并发开始时间
	goroutinetest(t)
	endTime := time.Now().UnixNano()
	useTime = float64(endTime-startTime) / 1e9 //记录所有请求完成时间
	fmt.Println("响应成功数:", ok)
	fmt.Println("相应失败数:", no)
	fmt.Println("qps :", fmt.Sprintf("%.4f", float64(num)/useTime))
}

func goroutinetest(t *testing.B) {
	for i := 0; i < num; i++ {
		go testLogin()
	}
	fmt.Println("all:", t.N)
	time.Sleep(time.Second * 1)
}

func testLogin() {
	data := url.Values{}
	//..
	data.Set("username", "admin")
	data.Set("password", "123456")
	resp, err := http.Post("http://127.0.0.1:8082/loginAuth", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		no += 1
	} else {
		ok += 1
	}
}
