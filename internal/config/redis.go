package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func NewRedis(config *viper.Viper, log *logrus.Logger) *redis.Client {
	address := config.GetString("redis.host") + ":" + strconv.Itoa(config.GetInt("redis.port"))
	password := config.GetString("redis.password")
	database := config.GetInt("redis.database")
	redisConnection := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	ctx := context.Background()

	pong, err := redisConnection.Ping(ctx).Result()
	if err != nil {
		log.Warnf("redis ping failed: %v", err)
	} else {
		log.Infof("redis connected : %v", pong)
	}

	return redisConnection
}
