package pool
//
//import (
//	"errors"
//	"fmt"
//	"sync"
//	"time"
//)
//
//var (
//	ErrMaxActiveConnReached = errors.New("MaxActiveConnReached")
//	ErrClosed = errors.New("pool is closed")
//)
//
//// Config 连接池相关配置
//type Config struct {
//	// 初始化连接数
//	InitialCap int
//	// 最大并发存活连接数
//	MaxCap int
//	// 最大空闲连接
//	MaxIdle int
//	// 生成连接的方法
//	Factory func() (interface{}, error)
//	// 关闭连接的方法
//	Close func(interface{}) error
//	// 检查连接是否有效的方法
//	Ping func(interface{}) error
//	// 连接最大空闲时间，超过该事件则将失效
//	IdleTimeout time.Duration
//}
//
//// channelPool 存放连接信息
//type channelPool struct {
//	mu                       sync.RWMutex
//	connects                 chan *idleConn
//	factory                  func() (interface{}, error)
//	close                    func(interface{}) error
//	ping                     func(interface{}) error
//	idleTimeout, waitTimeOut time.Duration
//	maxActive                int
//	openingConnects          int
//}
//
//type idleConn struct {
//	conn interface{}
//	t    time.Time
//}
//
//// Pool 基本方法
//type Pool interface {
//	Get() (interface{}, error)
//
//	Put(interface{}) error
//
//	Close(interface{}) error
//
//	Release()
//
//	Len() int
//}
//
//// Init 初始化连接
//func Init(poolConfig *Config) (Pool, error) {
//	if ! (poolConfig.InitialCap <= poolConfig.MaxIdle && poolConfig.MaxCap >= poolConfig.MaxIdle && poolConfig.InitialCap >= 0 ){
//		return nil, errors.New("invalid capacity settings")
//	}
//	if poolConfig.Factory == nil {
//		return nil, errors.New("invalid factory func settings")
//	}
//	if poolConfig.Close == nil {
//		return nil, errors.New("invalid close func settings")
//	}
//
//	c := &channelPool{
//		connects:        make(chan *idleConn, poolConfig.MaxIdle),
//		factory:      	 poolConfig.Factory,
//		close:        	 poolConfig.Close,
//		idleTimeout:  	 poolConfig.IdleTimeout,
//		maxActive:    	 poolConfig.MaxCap,
//		openingConnects: poolConfig.InitialCap,
//	}
//
//	if poolConfig.Ping != nil {
//		c.ping = poolConfig.Ping
//	}
//
//	for i := 0; i < poolConfig.InitialCap; i++ {
//		conn, err := c.factory()
//		if err != nil {
//			c.Release()
//			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
//		}
//		c.connects <- &idleConn{conn: conn, t: time.Now()}
//	}
//
//	return c, nil
//}
//
//// getConnects 获取所有连接
//func (c *channelPool) getConnects() chan *idleConn {
//	c.mu.Lock()
//	connects := c.connects
//	c.mu.Unlock()
//	return connects
//}
//
//// Get 从pool中取一个连接
//func (c *channelPool) Get() (interface{}, error) {
//	connects := c.getConnects()
//	if connects == nil {
//		return nil, ErrClosed
//	}
//	for {
//		select {
//		case wrapConn := <-connects:
//			if wrapConn == nil {
//				return nil, ErrClosed
//			}
//			// 判断是否超时，超时则丢弃
//			if timeout := c.idleTimeout; timeout > 0 {
//				if wrapConn.t.Add(timeout).Before(time.Now()) {
//					c.Close(wrapConn.conn)
//					continue
//				}
//			}
//			// 判断是否失效，失效则丢弃，如果用户没有设定 ping 方法，就不检查
//			if c.ping != nil {
//				if err := c.Ping(wrapConn.conn); err != nil {
//					c.Close(wrapConn.conn)
//					continue
//				}
//			}
//			return wrapConn.conn, nil
//		default:
//			c.mu.Lock()
//			defer c.mu.Unlock()
//			if c.openingConnects >= c.maxActive {
//				return nil, ErrMaxActiveConnReached
//			}
//			if c.factory == nil {
//				return nil, ErrClosed
//			}
//			conn, err := c.factory()
//			if err != nil {
//				return nil, err
//			}
//			c.openingConnects++
//			return conn, nil
//		}
//	}
//}
//
//// Put 将连接放回pool中
//func (c *channelPool) Put(conn interface{}) error {
//	if conn == nil {
//		return errors.New("connection is nil. rejecting")
//	}
//
//	c.mu.Lock()
//
//	if c.connects == nil {
//		c.mu.Unlock()
//		return c.Close(conn)
//	}
//
//	select {
//	case c.connects <- &idleConn{conn: conn, t: time.Now()}:
//		c.mu.Unlock()
//		return nil
//	default:
//		c.mu.Unlock()
//		// 连接池已满，直接关闭该连接
//		return c.Close(conn)
//	}
//}
//
//// Close 关闭单条连接
//func (c *channelPool) Close(conn interface{}) error {
//	if conn == nil {
//		return errors.New("connection is nil. rejecting")
//	}
//	c.mu.Lock()
//	defer c.mu.Unlock()
//	if c.close == nil {
//		return nil
//	}
//	c.openingConnects--
//	return c.close(conn)
//}
//
//// Ping 检查单条连接是否有效
//func (c *channelPool) Ping(conn interface{}) error {
//	if conn == nil {
//		return errors.New("connection is nil. rejecting")
//	}
//	return c.ping(conn)
//}
//
//// Release 释放连接池中所有连接
//func (c *channelPool) Release() {
//	c.mu.Lock()
//	connects := c.connects
//	c.connects = nil
//	c.factory = nil
//	c.ping = nil
//	closeFun := c.close
//	c.close = nil
//	c.mu.Unlock()
//
//	if connects == nil {
//		return
//	}
//
//	close(connects)
//	for wrapConn := range connects {
//		err := closeFun(wrapConn.conn)
//		if err != nil {
//			return
//		}
//	}
//}
//
//// Len 连接池中已有的连接
//func (c *channelPool) Len() int {
//	return len(c.getConnects())
//}