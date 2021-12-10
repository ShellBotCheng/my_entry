package entity

// GetUserInfoReq 获取用户信息请求体
type GetUserInfoReq struct {
	Username string
}

// GetUserInfoResp 返回体
type GetUserInfoResp struct {
	Status   int32
	Username string
	Nickname string
	Password string
	Salt     string
	PicUrl   string
}

// EditUserReq 更新用户信息请求体
type EditUserReq struct {
	Username string
	Nickname string
	PicUrl   string
}

// EditUserResp 返回体
type EditUserResp struct {
	Status int32
}

// GetSessionReq 获取会话信息
type GetSessionReq struct {
	SessionID string
}

// GetSessionResp 返回体
type GetSessionResp struct {
	Status      int32
	SessionInfo string
}

// SetSessionReq 设置会话信息 请求体
type SetSessionReq struct {
	SessionId   string
	SessionInfo string
}

type SetSessionResp struct {
	Status int32
}

// RefreshSessionReq 刷新会话
type RefreshSessionReq struct {
	SessionId string
}

type RefreshSessionResp struct {
	Status int32
}

// DelSessionReq 删除会话
type DelSessionReq struct {
	SessionId string
}

type DelSessionResp struct {
	Status int32
}
