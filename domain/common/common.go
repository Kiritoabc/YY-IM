package common

// LoginReq  登录请求
type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// LoginResp  登录响应
type LoginResp struct {
	UserID    int    `json:"userID"`
	UserName  string `json:"userName"`
	SessionId string `json:"sessionId"`
}
