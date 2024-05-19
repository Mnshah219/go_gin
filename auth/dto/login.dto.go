package dto

type LoginRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}
