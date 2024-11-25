package configs

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Mysql *MysqlConfig `json:"Mysql"`
	Redis *RedisConfig `json:"Redis"`
	JWT   *JWTConfig   `json:"JWT"`
}

type MysqlConfig struct {
	Host   string `json:"Host"`
	Port   int    `json:"Port"`
	User   string `json:"User"`
	Pass   string `json:"Pass"`
	DbName string `json:"DbName"`
}

type RedisConfig struct {
	Addr string `json:"Addr"`
	Pass string `json:"Pass"`
	DB   int    `json:"DB"`
}

type JWTConfig struct {
	Secret string `json:"Secret"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func Init(config *Config) {
	mc := config.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mc.User, mc.Pass, mc.Host, mc.Port, mc.DbName)

	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	rc := config.Redis
	_rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Pass,
		DB:       rc.DB,
	})

	rds = _rdb
}
