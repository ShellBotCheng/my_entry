package controller

import (
	"encoding/json"
	"io"
	HttpCnf "myEntry/httpServer/conf"
	"myEntry/pkg/content"
	"myEntry/pkg/log"
	"myEntry/pkg/utils"
	"net/http"
	"os"
	"strings"
)

type UploadInfo struct {
	Status int32  `json:"status"`
	PicUrl string `json:"pic_url"`
	Msg    string `json:"message"`
}

// Upload 上传
func Upload(w http.ResponseWriter, r *http.Request) {
	resp := UploadInfo{
		Status: content.SucCode,
	}

	// 检查会话
	//_, b := service.CheckSession(r)
	//if !b {
	//	resp.Msg = fmt.Sprintf("%s", content.SessionExpiredErrOR)
	//	msg, _ := json.Marshal(resp)
	//	_, _ = w.Write(msg)
	//	return
	//}
	picture, _, err := r.FormFile("file")
	if err != nil {
		return
	}
	picPath := HttpCnf.Conf.Http.PicPath + fileName(16)
	file, err := os.OpenFile(picPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Logger.Errorf("open file : %s", err)
		return
	}

	defer file.Close()

	_, err = io.Copy(file, picture)
	if err != nil {
		log.Error("err:%s", err)
		resp.Status = content.ServerError
		msg, _ := json.Marshal(resp)
		_, err = w.Write(msg)
		return
	}
	resp.PicUrl = strings.TrimLeft(picPath, ".")
	msg, _ := json.Marshal(resp)
	_, _ = w.Write(msg)
	return
}

// 更改获取文件名称

func fileName(len int) string {
	file, _ := utils.GenerateRandomString(len)
	return file + ".jpeg"
}
