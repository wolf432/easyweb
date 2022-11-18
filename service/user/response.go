package user

type ResponseUserInfo struct {
	Uid      uint   `json:"uid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
