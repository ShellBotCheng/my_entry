package service

import (
	"myEntry/pkg/content"
	"myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/lru"
	"myEntry/pkg/mysql"
	"myEntry/tcpServer/mapper"
	"myEntry/tcpServer/model"
)

func GetUserInfo(req entity.GetUserInfoReq) (resp entity.GetUserInfoResp, err error) {
	resp.Status = content.SucCode

	user := model.User{}
	// 查缓存
	cache, ok := lru.Client.Get(req.Username)
	if ok {
		user = cache.(model.User)
	} else {
		// 查询db
		user, err = mapper.GetUser(req.Username)
		if err != nil {
			resp.Status = content.TcpServerError
			return
		}
		lru.Client.Add(user.Username, user)
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
	lru.Client.Del(req.Username)
	tx.Commit()
	return
}
