package cache

import (
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

// UserLocation cache object
type UserLocation time.Location

var userConfigCache sync.Map
var redisClient *redis.Client

// InitCache initializes the cache
func InitCache() {
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

	log.Info().Msgf("initialized Cache. Redis is playing PING %v", pong)
}

// SetUserLocation stores the Location in the cache, and writes it to redis.
func SetUserLocation(userID string, location *time.Location) {
	userConfigCache.Store(userID, location)

	err := redisClient.HSet("user:"+userID+":config", "timezone", location.String()).Err()
	if err != nil {
		log.Error().Msgf("unable to store user timezone in redis %v", err.Error())
		return
	}
}

// GetUserLocation gets the Location in the cache, and if it wasn't found, it will
// check Redis and store the value found there, if any.
func GetUserLocation(userID string) *time.Location {
	// Check if we have it in the cache
	entry, ok := userConfigCache.Load(userID)
	if !ok {
		// wasn't in the cache; try redis
		tz, err := redisClient.HGet("user:"+userID+":config", "timezone").Result()
		if err != nil {
			// wasn't in redis either, but this is not an error.
			return nil
		}
		// Was in redis! Parse the Location
		location, err := time.LoadLocation(tz)
		if err != nil {
			log.Error().Msgf("unable to parse timezone string from redis: %v", err.Error())
			return nil
		}
		// and update the cache so we don't get into this mess again
		SetUserLocation(userID, location)
		return location
	}

	// found it in the cache!
	location, ok := entry.(*time.Location)
	if !ok {
		log.Error().Msgf("unexpected type %T found in cache", entry)
	}
	return location
}
