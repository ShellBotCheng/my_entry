package controller

import (
	"encoding/json"
	"fmt"
	"myEntry/httpServer/entity"
	"myEntry/httpServer/service"
	"myEntry/httpServer/view"
	"myEntry/pkg/content"
	"net/http"
	"strings"
)

// Login 登录页面
func Login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if r.RequestURI != "/" && r.RequestURI != "/login" {
		view.Render(data, w, "404")
		return
	}

	// 检查会话
	_, b := service.CheckSession(r)
	if b {
		http.Redirect(w, r, "/user/info", http.StatusFound)
	}
	view.Render(data, w, "login")
}

// LoginAuth 用户登录验证
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	resp := entity.UserAuth{
		Status: content.SucCode,
	}
	uname := r.FormValue("username")
	pwd := r.FormValue("password")

	if uname == "" || pwd == "" {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Status = content.ParamsMiss
		resp.Msg = "params miss"
		msg, _ := json.Marshal(resp)
		_, _ = w.Write(msg)
		return
	}
	// 登录验证服务
	sessionId, err := service.LoginAuth(uname, pwd)
	if err != nil {
		resp.Status = content.AccountError
		resp.Msg = fmt.Sprintf("%s", err)
		msg, _ := json.Marshal(resp)
		_, _ = w.Write(msg)
		return
	}
	cookie := &http.Cookie{
		Name:     "sessionId",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   int(content.SessionExpireTime),
	}
	http.SetCookie(w, cookie)
	// 服务端缓存session
	info := entity.SessionInfo{
		SessionId: sessionId,
		Username:  uname,
		IP:        getIp(r),
	}
	service.SetSession(sessionId, info)
	resp.Msg = "success"
	msg, _ := json.Marshal(resp)
	_, _ = w.Write(msg)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// 删除会话
	cookie, _ := r.Cookie("sessionID")
	service.DelSession(cookie.Value)
	// 重置cookie
	httpCookie := &http.Cookie{
		Name:   "sessionID",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, httpCookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func getIp(r *http.Request) (ip string) {
	addr := r.RemoteAddr
	arr := strings.Split(addr, ":")
	if len(arr) == 2 {
		ip = arr[0]
	}
	return
}
