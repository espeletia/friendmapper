package domain

type UserData struct {
	Username    string
	Email       string
	DisplayName string
	Hash        string
}

type User struct {
	ID             int64
	DisplayName    string
	Email          string
	Username       string
	HashedPassword string
}

type LoginCreds struct {
	Email    string
	Password string
}
