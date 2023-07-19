package router

import (
	"web-server/internal/authen"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RouterEngine() *gin.Engine {
	r := gin.Default()

	r.Use(CORS())
	r.POST("/register", authen.Register)
	r.POST("/login", authen.Login)
	user := r.Group("/user")
	user.Use(authen.AuthenUser)
	{
		user.GET("/username",authen.GetUsername)
		user.POST("/logout", authen.Logout)
	}

	return r
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		targets := viper.GetString("cors.target")
		c.Writer.Header().Set("Access-Control-Allow-Origin", targets)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
