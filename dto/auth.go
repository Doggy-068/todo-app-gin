package dto

type AuthLoginPostDto struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthLoginPostReturnDto struct {
	ApiKey string `json:"api_key" binding:"required"`
}
