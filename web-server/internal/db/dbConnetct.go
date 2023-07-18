package db

import (
	"context"
	"database/sql"
	"time"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var database *sql.DB
var redisDB *redis.Client

func ConnectDB() {

	address := viper.GetString("connection.dbURL")
	db, err := sql.Open("mysql", address)
	if err != nil {
		panic("Cannot access mariadb server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	database = db

}
func DisconnectDB() {
	database.Close()
}

func ConnectRedis() {
	address := viper.GetString("connection.redisURL")
	// secret := viper.GetString("connection.redisSecret")

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0, // use default database
	})

	// Test the connection to Redis
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	redisDB = client
}
func DisconnectRedis() {
	redisDB.Close()
}

func GetDatabase() *sql.DB {
	return database
}
func GetRedis() *redis.Client {
	return redisDB
}
