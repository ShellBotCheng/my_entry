package service

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"myEntry/httpServer/common"
	"myEntry/httpServer/entity"
	"myEntry/pkg/content"
	tcpEntity "myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/rpc"
	"myEntry/pkg/utils"
	"net"
	"net/http"
	"strings"
)

func init() {
	gob.Register(tcpEntity.GetSessionReq{})
	gob.Register(tcpEntity.GetSessionResp{})
	gob.Register(tcpEntity.RefreshSessionReq{})
	gob.Register(tcpEntity.RefreshSessionResp{})
	gob.Register(tcpEntity.SetSessionReq{})
	gob.Register(tcpEntity.SetSessionResp{})
	gob.Register(tcpEntity.DelSessionReq{})
	gob.Register(tcpEntity.DelSessionResp{})
}

// createSessionID 生成sessionID，随机字符串+用户名的MD5值
func createSessionID(uname string) string {
	randomString, _ := utils.GenerateRandomString(utils.RandomLen)
	s := fmt.Sprintf("%s.%s", randomString, uname)
	sessionID := utils.MD5(s)
	return sessionID
}

// CheckSession 检查会话
func CheckSession(r *http.Request) (sessionInfo entity.SessionInfo, ok bool) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Error("get cookie err:%s", err)
		return
	}
	sessionID := cookie.Value
	sessionInfo, err = GetSession(sessionID)
	if err != nil {
		return
	}

	arr := strings.Split(r.RemoteAddr, ":")
	if len(arr) != 2 || arr[0] == "" {
		return
	}

	if sessionInfo.IP != arr[0] {
		return
	}
	// 更新会话时间
	refreshSessionRpc(sessionID)
	return sessionInfo, true
}

// SetSession 设置会话
func SetSession(sessionID string, info entity.SessionInfo) bool {
	conn, err := common.TcpPool.Get()
	if err != nil {
		log.Error("get conn err:%s", err)
		return false
	}

	cli := rpc.NewClient(conn.(net.Conn))

	cli.CallRPC("SetSession", &rpc.SetSession)
	infoStr, _ := json.Marshal(info)
	req := tcpEntity.SetSessionReq{
		SessionId:   sessionID,
		SessionInfo: string(infoStr),
	}
	r := rpc.SetSession(req)
	if err != nil {
		log.Error("SetSession Error:%s", err)
		return false
	}
	if r.Status != content.SucCode {
		log.Error("SetSession Status:%d", r.Status)
		return false
	}
	log.Info("set session ok")

	_ = common.TcpPool.Put(conn)
	return true
}

// GetSession 获取会话信息
func GetSession(sessionID string) (sessionInfo entity.SessionInfo, err error) {
	sessionInfo = entity.SessionInfo{}
	conn, err := common.TcpPool.Get()
	defer common.TcpPool.Put(conn)
	if err != nil {
		log.Error("get conn err:%s", err)
		return
	}

	cli := rpc.NewClient(conn.(net.Conn))

	cli.CallRPC("GetSession", &rpc.GetSession)
	req := tcpEntity.GetSessionReq{
		SessionID: sessionID,
	}

	r := rpc.GetSession(req)
	if err != nil {
		log.Error("GetSession Error:%s", err)
		return
	}
	if r.Status != content.SucCode {
		log.Error("GetSession Status:%d", r.Status)
		return
	}

	_ = json.Unmarshal([]byte(r.SessionInfo), &sessionInfo)
	sessionInfo.SessionId = sessionID
	return sessionInfo, nil
}

// refreshSessionRpc 刷新会话
func refreshSessionRpc(sessionID string) {
	conn, err := common.TcpPool.Get()
	defer common.TcpPool.Put(conn)
	if err != nil {
		log.Error("get conn err:%s", err)
		return
	}

	cli := rpc.NewClient(conn.(net.Conn))

	cli.CallRPC("RefreshSession", &rpc.RefreshSession)
	req := tcpEntity.RefreshSessionReq{
		SessionId: sessionID,
	}

	r := rpc.RefreshSession(req)
	if err != nil {
		log.Error("RefreshSession Error:%s", err)
		return
	}
	if r.Status != content.SucCode {
		log.Error("RefreshSession Status:%d", r.Status)
		return
	}
	log.Info("refresh session success")

}

// DelSession 销毁会话
func DelSession(sessionID string) {
	conn, err := common.TcpPool.Get()
	defer common.TcpPool.Put(conn)
	if err != nil {
		log.Error("get conn err:%s", err)
		return
	}

	cli := rpc.NewClient(conn.(net.Conn))

	cli.CallRPC("DelSession", &rpc.DelSession)
	req := tcpEntity.DelSessionReq{
		SessionId: sessionID,
	}

	r := rpc.DelSession(req)
	if err != nil {
		log.Error("err:%s", err)
		return
	}
	if r.Status != content.SucCode {
		log.Error("err ret:%d", r.Status)
		return
	}
	log.Info("refresh session ok")
}
