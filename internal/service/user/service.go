package user

import (
	"ccgin/internal/models"
	"ccgin/pkgs/jwtool"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const cacheKeyByUserID = "user:id:%d"
const cacheKeyByUsername = "user:username:%s"

type Service interface {
	i()
	Login(ctx context.Context, username, password string) (*LoginResult, error)
	Check(ctx context.Context, token string) (*models.User, error)

	Info(ctx context.Context, id int) (*models.User, error)
	InfoByUsername(ctx context.Context, username string) (*models.User, error)
}

type service struct {
	db      *gorm.DB
	rds     *redis.Client
	secret  string
	jwtTool *jwtool.JWTool
}

func New(db *gorm.DB, redis *redis.Client, secret string) Service {
	return &service{
		db:      db,
		rds:     redis,
		secret:  secret,
		jwtTool: jwtool.New(secret),
	}
}
func (s *service) i() {}
