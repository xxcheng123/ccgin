package user

import (
	"ccgin/internal/codes"
	"ccgin/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type LoginResult struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
	ExpireIn   int64  `json:"expireIn"`
}

func (s *service) Login(ctx context.Context, username, password string) (*LoginResult, error) {
	var user = &models.User{}

	err := s.db.WithContext(ctx).Where("username = ?", username).First(user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, codes.UserNotFound
		}

		return nil, err
	}

	if user.Password != password {
		return nil, codes.UserPasswordErr
	}

	if !user.StatusOk() {
		return nil, codes.AccountError
	}

	token, err := s.jwtTool.Generate(user.ID, user.Username, user.Version, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	now := time.Now().Add(24 * time.Hour)

	return &LoginResult{
		Token:      token,
		ExpireTime: now.Unix(),
		ExpireIn:   24 * 60 * 60,
	}, nil
}
