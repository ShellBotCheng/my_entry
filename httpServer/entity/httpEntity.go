package entity

// UserInfo 用户信息
type UserInfo struct {
	Username string // 用户名
	Nickname string // 昵称
	PicUrl   string // 头像
	Password string // 密码
	Salt     string // 盐
}

// SessionInfo 会话信息
type SessionInfo struct {
	Username  string `json:"username"`  // 用户名
	IP        string `json:"ip"`        // 用户IP地址
	SessionId string `json:"sessionId"` // 会话ID
}

// UserAuth 登录验证实体
type UserAuth struct {
	SessionId string `json:"sessionId"`
	Status    int32  `json:"status"`
	Msg       string `json:"message"`
}

// UserEdit 用户编辑实体
type UserEdit struct {
	Status   int32  `json:"status"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	PicUrl   string `json:"pic_url"`
	Msg      string `json:"message"`
}
