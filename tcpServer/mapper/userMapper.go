package mapper

import (
	"myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/mysql"
	"myEntry/tcpServer/model"
)

func GetUser(username string) (user model.User, err error) {
	row := mysql.Db.QueryRow("select username, password, nickname, salt, pic_url from user where username=?", username)
	err = row.Scan(&user.Username, &user.Password, &user.Nickname, &user.Salt, &user.PicUrl)
	if err != nil {
		log.Error("GetUser Error:%s", err)
		return
	}
	return
}

func UpdateUser(q entity.EditUserReq) (affected int, err error) {
	affected = 0
	res, err := mysql.Db.Exec("update user set nickname=?, pic_url=? where username=?", q.Nickname, q.PicUrl, q.Username)
	if err != nil {
		log.Error("user %d update err %v", res, err)
		return
	}
	affected64, err := res.RowsAffected()
	if err != nil {
		log.Error("user %d update err %v", affected64, err)
		return
	}
	affected = int(affected64)
	return
}
