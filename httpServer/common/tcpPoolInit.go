package common

import (
	"myEntry/pkg/log"
	"myEntry/pkg/pool"
	"net"
	"time"
)

var (
	TcpPool pool.Pool
)

type TcpConfig struct {
	NetWork string
	// 链接地址
	Addr string
	// 连接池中拥有的最小连接数
	InitialCap int
	// 最大并发存活连接数
	MaxCap int
	// 最大空闲连接
	MaxIdle int
	// 生成连接的方法
	Factory func() (interface{}, error)
	// 关闭连接的方法
	Close func(interface{}) error
	// 检查连接是否有效的方法
	Ping func(interface{}) error
	// 连接最大空闲时间，超过该事件则将失效
	IdleTimeout time.Duration
}

// TcpPoolInit 初始化Tcp连接池
func TcpPoolInit(c *TcpConfig) error {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return net.Dial(c.NetWork, c.Addr) }
	//f 关闭连接的方法
	f := func(v interface{}) error { return v.(net.Conn).Close() }
	//创建一个连接池
	poolConfig := &pool.Config{
		InitialCap: c.InitialCap,
		MaxIdle:    c.MaxIdle,
		MaxCap:     c.MaxCap,
		Factory:    factory,
		Close:      f,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := pool.Init(poolConfig)
	if err != nil {
		log.Error("TcpPoolInit error: ", err)
		return err
	}
	TcpPool = p
	//查看当前连接中的数量
	current := TcpPool.Len()
	log.Info("TcpPool Len= ", current)
	return nil
}
