package dto

type AuthRequest struct {
	EmailAddress string `json:"email_address" binding:"required"`
	Password string `json:"password" binding:"required"`
}


type Token struct {
	Token string `json:"token" binding:"required"`
}