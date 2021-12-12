package service

import (
	"encoding/gob"
	"errors"
	"myEntry/httpServer/common"
	"myEntry/httpServer/entity"
	"myEntry/pkg/content"
	tcpEntity "myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/rpc"
	"net"
	"net/http"
)

func init() {
	gob.Register(tcpEntity.GetUserInfoReq{})
	gob.Register(tcpEntity.GetUserInfoResp{})
	gob.Register(tcpEntity.EditUserReq{})
	gob.Register(tcpEntity.EditUserResp{})
}

var UserNotExistErr = errors.New("user not exist")

// GetUserInfo 获取用户信息
func GetUserInfo(uname string) (user entity.UserInfo, err error) {
	conn, err := common.TcpPool.Get()
	if err != nil {
		log.Error("get conn err:%s", err)
		return
	}
	//conn, err := net.Dial("tcp", "127.0.0.1:8081")
	cli := rpc.NewClient(conn.(net.Conn))
	cli.CallRPC("GetUser", &rpc.GetUser)
	req := tcpEntity.GetUserInfoReq{
		Username: uname,
	}
	defer common.TcpPool.Put(conn)
	r, err := rpc.GetUser(req)
	if err != nil {
		log.Error("GetUserInfo Error:%s", err)
		return
	}
	if r.Status != content.SucCode {
		log.Error("GetUserInfo Status:%d", r.Status)
		return
	}
	user.Username = r.Username
	user.Nickname = r.Nickname
	user.PicUrl = r.PicUrl
	user.Password = r.Password
	user.Salt = r.Salt
	return user, nil
}

//UpdateUserInfo 修改用户信息
func UpdateUserInfo(r *http.Request) (user entity.UserInfo, err error) {
	conn, err := common.TcpPool.Get()
	if err != nil {
		log.Error("get conn err:%s", err)
		return
	}
	//conn, err = net.Dial("tcp", "127.0.0.1:8081")
	cli := rpc.NewClient(conn.(net.Conn))
	cli.CallRPC("EditUser", &rpc.EditUser)
	req := tcpEntity.EditUserReq{
		Username: r.FormValue("username"),
		Nickname: r.FormValue("nickname"),
		PicUrl:   r.FormValue("picUrl"),
	}
	defer common.TcpPool.Put(conn)
	resp, err := rpc.EditUser(req)
	if err != nil {
		log.Error("EditUser Error:%s", err)
		return
	}
	if resp.Status != content.SucCode {
		log.Error("EditUser Status:%d", resp.Status)
		return
	}
	user.Username = req.Username
	user.Nickname = req.Nickname
	user.PicUrl = req.PicUrl
	return user, nil
}
