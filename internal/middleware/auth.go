package middleware

import (
	"ccgin/internal/codes"
	"ccgin/internal/service/user"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strings"
)

func Auth(db *gorm.DB, rds *redis.Client, secret string) gin.HandlerFunc {
	userService := user.New(db, rds, secret)
	return func(ctx *gin.Context) {
		tk := ctx.GetHeader("Authorization")
		tk = strings.TrimPrefix(tk, "Bearer ")
		u, err := userService.Check(ctx, tk)
		if err != nil {
			ctx.Abort()
			if errors.Is(err, jwt.ErrTokenExpired) {
				codes.UserAuthExpired.EmptyResponse(ctx)

				return
			}

			codes.UserJWTErr.EmptyResponse(ctx)

			return
		}

		ctx.Set("userId", u.ID)
		ctx.Set("username", u.Username)
		ctx.Set("user", u)
	}
}
