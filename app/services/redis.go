package services

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// REDIS null if not initialized
var REDIS *redis.Client

// InitRedis initialize redis connection
func InitRedis() {
	if nil == REDIS {
		redisHost := viper.GetString("REDIS_HOST")
		redisPort := viper.GetString("REDIS_PORT")
		if redisHost != "" {
			REDIS = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
				Password: viper.GetString("REDIS_PASS"),
				DB:       viper.GetInt("REDIS_INDEX"),
			})
		}
	}
}
