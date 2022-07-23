package models

type Login struct {
	Password string `json:"password" binding:"required" example:"password"`
	Email    string `json:"email" binding:"required" example:"user1@mail.com"`
}

type TokenDetails struct {
	ID        uint   `json:"id" example:"123"`
	Token     string `json:"token"  example:"auth-jwt-token"`
	AtExpires int64  `json:"atExpires"  example:"1658548537"`
}
