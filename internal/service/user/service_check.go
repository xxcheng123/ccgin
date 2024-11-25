package user

import (
	"ccgin/internal/models"
	"context"
)

func (s *service) Check(ctx context.Context, token string) (*models.User, error) {
	u, err := s.jwtTool.Parse(token)
	if err != nil {
		return nil, err
	}

	user, err := s.Info(ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	if !user.StatusOk() {
		return nil, err
	}

	if user.Version > u.Version {
		return nil, err
	}

	return user, nil
}
