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

func GetUserInfo(req entity.GetUserInfoReq) (resp entity.GetUserInfoResp, err error) {
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
		redis.SetEx(user.Username, string(userJson), content.SessionExpTime)
	}
	// 装载数据
	resp.Username = user.Username
	resp.Nickname = user.Nickname
	resp.PicUrl = user.PicUrl
	resp.Salt = user.Salt
	resp.Password = user.Password
	return
}

func UpdateUserInfo(req entity.EditUserReq) (resp entity.EditUserResp, err error) {
	resp.Status = content.SucCode

	tx, err := mysql.Db.Begin() // 开启事务
	defer tx.Rollback()

	if err != nil {
		resp.Status = content.TcpServerError
		log.Error("begin trans failed, err:%v\n", err)
		return
	}
	n, err := mapper.UpdateUser(req)
	if err != nil {
		resp.Status = content.TcpServerError
		log.Error("UpdateUser failed, err:%v\n", err)
		return
	}
	if n == 0 {
		return
	}
	// 清除缓存
	_, err = redis.Del(req.Username)
	if err != nil {
		resp.Status = content.TcpServerError
	}
	tx.Commit()
	return
}
