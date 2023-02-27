package mapper

import (
	"fmt"
	"hash/crc32"
	"myEntry/pkg/entity"
	"myEntry/pkg/log"
	"myEntry/pkg/mysql"
	"myEntry/tcpServer/model"
	"strconv"
)

func GetUser(username string) (user model.User, err error) {
	table := getTableName(username)
	row := mysql.GetDB().QueryRow(fmt.Sprintf("select username, password, nickname, salt, pic_url from %s where username='%s'",  table, username))
	err = row.Scan(&user.Username, &user.Password, &user.Nickname, &user.Salt, &user.PicUrl)
	if err != nil {
		log.Error("GetUser Error:%s:%s:%s", err,table,username)
		return
	}
	return
}

func UpdateUser(q entity.EditUserReq) (affected int, err error) {
	affected = 0
	table := getTableName(q.Username)

	res, err := mysql.Db.Exec(fmt.Sprintf("update %s set nickname='%s', pic_url='%s' where username='%s'",table, q.Nickname, q.PicUrl, q.Username))
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

// getTableName 获取表名
func getTableName(uname string) string {
	return fmt.Sprintf("user_%04s", strconv.Itoa(int(crc32.ChecksumIEEE([]byte(uname))%128)))
}
