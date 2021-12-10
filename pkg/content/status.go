package content

const (
	// SucCode Common code
	SucCode int32 = 200 //成功码

	// http error status

	ParamsMiss    int32 = 1001 // 参数缺失
	AccountError  int32 = 1002 // 用户信息异常
	ServerError   int32 = 1003 // 服务内部错误
	ExpireSession int32 = 1004 // session过期

	// tcp error status

	TcpAccountError int32 = 2002 // 用户信息异常
	TcpServerError  int32 = 2003 // 服务内部错误
	TcpSessionMiss  int32 = 2004 // session过期
)