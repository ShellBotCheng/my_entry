package service

import (
	"myEntry/pkg/content"
	"myEntry/pkg/entity"
	"myEntry/pkg/redis"
)

// GetSession 获取会话信息
func GetSession(req entity.GetSessionReq ) entity.GetSessionResp {
	resp := entity.GetSessionResp{
		Status : content.SucCode,
	}
	session, err := redis.Get(req.SessionID)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	if session == content.EmptyString{
		resp.Status = content.TcpSessionMiss
	}
	resp.SessionInfo = session
	return resp
}

// SetSession  设置会话信息
func SetSession(req entity.SetSessionReq) entity.SetSessionResp {
	resp := entity.SetSessionResp{
		Status : content.SucCode,
	}

	err := redis.SetEx(req.SessionId, req.SessionInfo, content.SessionExpireTime)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	return resp
}

// RefreshSession 刷新会话
func RefreshSession(req entity.RefreshSessionReq ) entity.RefreshSessionResp{
	resp := entity.RefreshSessionResp{
		Status : content.SucCode,
	}

	b, err := redis.Refresh(req.SessionId, content.SessionExpireTime)
	if err != nil || !b{
		resp.Status = content.TcpServerError
	}
	return resp
}

// DelSession 移除会话
func DelSession(req entity.DelSessionReq ) entity.DelSessionResp{
	resp := entity.DelSessionResp{
		Status : content.SucCode,
	}

	_, err := redis.Del(req.SessionId)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	return resp
}

