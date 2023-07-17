package main

import (
	"web-server/internal/db"
	"web-server/internal/util"
	"web-server/internal/router"
	"github.com/spf13/viper"
	"fmt"
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
