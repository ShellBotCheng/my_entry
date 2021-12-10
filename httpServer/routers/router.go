package routers

import (
	"myEntry/httpServer/controller"
	"net/http"
)

func InitRouter() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", controller.Login)
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/loginAuth", controller.LoginAuth)
	mux.HandleFunc("/logout", controller.Logout)
	mux.HandleFunc("/userInfo", controller.UserInfo)
	mux.HandleFunc("/upload", controller.Upload)
	mux.HandleFunc("/userEdit", controller.Edit)
	// 静态资源路径
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}
