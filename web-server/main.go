package main

import (
	"fmt"
	"web-server/internal/db"
	"web-server/internal/router"
	"web-server/internal/util"

	"github.com/spf13/viper"
)

func main() {
	util.InitViper()
	db.ConnectDB()
	defer db.DisconnectDB()

	db.ConnectRedis()
	defer db.DisconnectRedis()

	r := router.RouterEngine()

	r.Run(fmt.Sprintf(":%d", viper.GetInt("connection.appPort")))
}
