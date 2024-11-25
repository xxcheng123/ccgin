package user

import (
	"ccgin/internal/codes"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *handler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req loginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			codes.ParamErr.Empty(ctx)

			return
		}

		result, err := h.userService.Login(ctx, req.Username, req.Password)

		if err != nil {
			if code, ok := codes.As(err); ok {
				code.Response(ctx, result)

				return
			}

			codes.DatabaseErr.Response(ctx, result)

			return
		}

		codes.Success.Response(ctx, result)
	}
}
