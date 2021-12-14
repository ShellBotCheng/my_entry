package service

import (
	"myEntry/pkg/content"
	"myEntry/pkg/entity"
	"myEntry/pkg/redis"
)

// GetSession 获取会话信息
func GetSession(req entity.GetSessionReq) (resp entity.GetSessionResp, err error) {
	resp.Status = content.SucCode

	session := redis.Get(req.SessionId)
	if session == content.EmptyString {
		resp.Status = content.TcpSessionMiss
		return
	}
	resp.SessionInfo = session
	return
}

// SetSession  设置会话信息
func SetSession(req entity.SetSessionReq) (resp entity.SetSessionResp, err error) {
	resp.Status = content.SucCode

	err = redis.SetEx(req.SessionId, req.SessionInfo, content.SessionExpTime)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	return
}

// RefreshSession 刷新会话
func RefreshSession(req entity.RefreshSessionReq) (resp entity.RefreshSessionResp, err error) {
	resp.Status = content.SucCode

	b, err := redis.Refresh(req.SessionId, content.SessionExpTime)
	if err != nil || !b {
		resp.Status = content.TcpServerError
	}
	return
}

// DelSession 移除会话
func DelSession(req entity.DelSessionReq) (resp entity.DelSessionResp, err error) {
	resp.Status = content.SucCode

	_, err = redis.Del(req.SessionId)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	return
}
