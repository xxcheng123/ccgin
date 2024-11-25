package user

import (
	"ccgin/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func (s *service) InfoByUsername(ctx context.Context, username string) (*models.User, error) {
	var key = fmt.Sprintf(cacheKeyByUsername, username)
	str, err := s.rds.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			var user models.User

			err = s.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
			if err != nil {
				return nil, err
			}

			data, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}

			s.rds.Set(ctx, key, string(data), 24*time.Hour)

			return &user, nil
		}

		return nil, err
	}

	var user models.User

	err = json.Unmarshal([]byte(str), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) Info(ctx context.Context, id int) (*models.User, error) {
	var key = fmt.Sprintf(cacheKeyByUserID, id)
	str, err := s.rds.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			var user models.User

			err = s.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

			if err != nil {
				return nil, err
			}

			data, err := json.Marshal(user)

			if err != nil {
				return nil, err
			}

			s.rds.Set(ctx, key, string(data), 24*time.Hour)

			return &user, nil
		}

		return nil, err
	}

	var user models.User

	err = json.Unmarshal([]byte(str), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
