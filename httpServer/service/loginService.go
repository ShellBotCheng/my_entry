package service

import (
	"fmt"
	"myEntry/pkg/content"
	"myEntry/pkg/log"
	"myEntry/pkg/utils"
)

// encryptedByMD5 获取pwd密文
func encryptedByMD5(pwd string, salt string) string {
	s := fmt.Sprintf("%s%s", pwd, salt)
	return utils.MD5(s)
}

// LoginAuth 用户登入
func LoginAuth(uname string, pwd string) (sessionId string, err error) {
	userInfo, err := GetUserInfo(uname)
	if err != nil {
		log.Error("GetUserInfo err:%s   uname:%s\n", err, uname)
		return
	}

	md5Pwd := encryptedByMD5(pwd, userInfo.Salt)
	log.Info("pwd:", md5Pwd)
	if userInfo.Password != md5Pwd {
		err = content.UserPwdError
		log.Error("LoginAuth err:%s   uname:%s", err, uname)
		return
	}
	sessionId = createSessionId(uname)
	return sessionId, nil
}
