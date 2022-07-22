package models

type Login struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TokenDetails struct {
	ID        uint   `json:"id"`
	Token     string `json:"token"`
	AtExpires int64  `json:"atExpires"`
}

type AccessDetails struct {
	UserId uint
}
