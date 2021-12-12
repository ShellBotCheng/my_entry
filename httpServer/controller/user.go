package controller

import (
	"encoding/json"
	"myEntry/httpServer/entity"
	"myEntry/httpServer/service"
	"myEntry/httpServer/view"
	"myEntry/pkg/content"
	"net/http"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	sessionInfo, b := service.CheckSession(r)
	if !b {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	user, _ := service.GetUserInfo(sessionInfo.Username)
	data["user"] = entity.UserInfo{
		Username: user.Username,
		Nickname: user.Nickname,
		PicUrl:   user.PicUrl,
	}
	view.Render(data, w, "user")
}

func Edit(w http.ResponseWriter, r *http.Request) {
	resp := entity.UserEdit{
		Status: content.SucCode,
	}
	uname := r.FormValue("username")
	nickname := r.FormValue("nickname")
	picUrl := r.FormValue("picUrl")

	if uname == "" || nickname == "" || picUrl == "" {
		resp.Status = content.ParamsMiss
		resp.Msg = "params miss"
		msg, _ := json.Marshal(resp)
		_, _ = w.Write(msg)
		return
	}
	// 检查会话
	_, b := service.CheckSession(r)
	if !b {
		resp.Status = content.ExpireSession
		resp.Msg = "Login has been Expired, please relogin"
		msg, _ := json.Marshal(resp)
		_, _ = w.Write(msg)
		return
	}

	user, err := service.UpdateUserInfo(r)
	if err != nil {
		resp.Status = content.ServerError
		resp.Msg = "Edit User Error"
		msg, _ := json.Marshal(resp)
		_, _ = w.Write(msg)
		return
	}
	resp.Username = user.Username
	resp.Nickname = user.Nickname
	resp.PicUrl = user.PicUrl
	msg, _ := json.Marshal(resp)
	_, _ = w.Write(msg)
	return
}
