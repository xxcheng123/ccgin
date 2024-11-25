package user

import (
	"ccgin/internal/codes"
	"github.com/gin-gonic/gin"
)

func (h *handler) Info() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := ctx.Get("userId")

		uid, _ := userId.(int)

		result, err := h.userService.Info(ctx, uid)

		if err != nil {
			codes.DatabaseErr.Response(ctx, err.Error())

			return
		}

		codes.Success.Response(ctx, result)
	}
}
