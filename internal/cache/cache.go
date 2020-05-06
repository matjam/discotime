package cache

import (
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

// UserConfig cache object
type UserConfig struct {
	TimeZone time.Location
}

var userConfigCache map[string]UserConfig
var redisClient *redis.Client

// InitCache initializes the cache
func InitCache() {
	userConfigCache = make(map[string]UserConfig)

	url := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(url)
	if err != nil {
		log.Error().Msgf("unable to parse REDIS_URL: %v", err.Error())
		return
	}
	redisClient = redis.NewClient(opt)

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Error().Msgf("unable to ping redis server %v", err.Error())
		return
	}

	log.Info().Msgf("%v", pong)
}

func SetUserConfig(username string, config UserConfig) {

}
