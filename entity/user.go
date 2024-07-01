package entity

type User struct {
	UserId   string `json:"userid"`
	UserName string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isadmin"`
}
