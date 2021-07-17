// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserReply
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisteResponse struct {
	UserReply
}

type UserinfoRequest struct {
	Userid string `json:"userid"`
	Token  string `json:"token"`
}

type UserinfoResponse struct {
	UserReply
}

type UserReply struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	JwtToken
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}
