package service

import (
	"encoding/json"
	"myEntry/pkg/content"
	"myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/mysql"
	"myEntry/pkg/redis"
	"myEntry/tcpServer/mapper"
)

func GetUserInfo(req entity.GetUserInfoReq) (resp entity.GetUserInfoResp) {
	resp.Status = content.SucCode

	// 查缓存
	cache, err := redis.Get(req.Username)
	if cache != content.EmptyString {
		err = json.Unmarshal([]byte(cache), &resp)
		if err == nil {
			return
		}
	}
	// 查询db
	user, err := mapper.GetUser(req.Username)
	if err != nil {
		resp.Status = content.TcpServerError
		return
	}
	// 序列化
	userJson, err := json.Marshal(user)
	if err == nil {
		redis.Set(user.Username, string(userJson))
	}
	// 装载数据
	resp.Username = user.Username
	resp.Nickname = user.Nickname
	resp.PicUrl = user.PicUrl
	resp.Salt = user.Salt
	resp.Password = user.Password
	return
}

func UpdateUserInfo(req entity.EditUserReq) entity.EditUserResp {
	resp := entity.EditUserResp{
		Status: content.SucCode,
	}
	tx, err := mysql.Db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			err := tx.Rollback()
			if err != nil {
				return resp
			} // 回滚
		}
		log.Error("begin trans failed, err:%v\n", err)
		return resp
	}
	n, err := mapper.UpdateUser(req)
	if err != nil {
		resp.Status = content.TcpServerError
		if tx != nil {
			err := tx.Rollback()
			if err != nil {
				return resp
			} // 回滚
		}
		return resp
	}
	if n == 0 {
		resp.Status = content.TcpAccountError
		return resp
	}
	// 清除缓存
	_, err = redis.Del(req.Username)
	if err != nil {
		resp.Status = content.TcpServerError
		if tx != nil {
			err := tx.Rollback()
			if err != nil {
				return resp
			} // 回滚
		}
	}
	return resp
}
