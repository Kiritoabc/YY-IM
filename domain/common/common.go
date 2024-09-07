package common

// LoginReq  登录请求
type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// LoginResp  登录响应
type LoginResp struct {
	Code      int    `json:"code"`
	UserID    int    `json:"userID"`
	UserName  string `json:"userName"`
	SessionId string `json:"sessionId"`
}

// SocketMsg socket消息
type SocketMsg struct {
	Action       string `json:"action"`
	TargetUserId int    `json:"targetUserId"`
	Message      string `json:"message"`
}
