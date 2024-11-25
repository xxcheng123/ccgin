package user

import (
	"ccgin/configs"
	"ccgin/internal/service/user"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Login() gin.HandlerFunc
	Info() gin.HandlerFunc
}

type handler struct {
	userService user.Service
}

func New() Handler {
	return &handler{
		userService: user.New(configs.DB(), configs.Rds(), configs.GetConfig().JWT.Secret),
	}
}
