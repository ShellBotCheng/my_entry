package model

type User struct {
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	Salt        string `json:"salt"`
	PicUrl      string `json:"pic_url"`
}
