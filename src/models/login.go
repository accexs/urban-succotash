package models

type Login struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type TokenDetails struct {
	Id        uint   `json:"id"`
	Token     string `json:"token"`
	AtExpires int64  `json:"atExpires"`
}

type AccessDetails struct {
	UserId uint
}
