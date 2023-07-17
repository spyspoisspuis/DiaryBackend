package dstore


import (
	"time"
	"web-server/internal/db"
)

const (
	OneDay = 86400
)

func LoginSession(username string, refreshToken string, expiration time.Duration) error {
	redisClient := db.GetRedis()
	return redisClient.Set(username, refreshToken, expiration).Err()
}

func GetToken(username string) (string, error) {
	redisClient := db.GetRedis()
    return redisClient.Get(username).Result()
}
func RemoveToken(username string) error {
	redisClient := db.GetRedis()
	return redisClient.Del(username).Err()
}


