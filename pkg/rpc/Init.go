package rpc

import (
	"encoding/gob"
	"myEntry/pkg/entity"
	"myEntry/tcpServer/service"
)

// rpc接口参数定义

var GetUser func(req entity.GetUserInfoReq) (entity.GetUserInfoResp, error)
var EditUser func(req entity.EditUserReq) (entity.EditUserResp, error)
var GetSession func(req entity.GetSessionReq) (entity.GetSessionResp, error)
var RefreshSession func(req entity.RefreshSessionReq) (entity.RefreshSessionResp, error)
var SetSession func(req entity.SetSessionReq) (entity.SetSessionResp, error)
var DelSession func(req entity.DelSessionReq) (entity.DelSessionResp, error)

// registerRpcEntity 注册实体
func registerRpcEntity() {
	gob.Register(entity.GetUserInfoReq{})
	gob.Register(entity.GetUserInfoResp{})
	gob.Register(entity.EditUserReq{})
	gob.Register(entity.EditUserResp{})
	gob.Register(entity.GetSessionReq{})
	gob.Register(entity.GetSessionResp{})
	gob.Register(entity.RefreshSessionReq{})
	gob.Register(entity.RefreshSessionResp{})
	gob.Register(entity.SetSessionReq{})
	gob.Register(entity.SetSessionResp{})
	gob.Register(entity.DelSessionReq{})
	gob.Register(entity.DelSessionResp{})
}

func InitRpcServer(addr string) *Server {
	registerRpcEntity()
	server := NewServer(addr)
	server.Register("GetSession", service.GetSession)
	server.Register("SetSession", service.SetSession)
	server.Register("RefreshSession", service.RefreshSession)
	server.Register("DelSession", service.DelSession)
	server.Register("GetUser", service.GetUserInfo)
	server.Register("EditUser", service.UpdateUserInfo)
	return server
}
