package configs

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

var rds *redis.Client

func Rds() *redis.Client {
	return rds
}
