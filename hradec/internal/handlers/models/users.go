package models

type UserData struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

type User struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
}

type LoginCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}
