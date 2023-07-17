package authen

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func AuthenUser(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	err := authorization(token)
	if err == ErrNoSession {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "UnAuthorized"})
		return
	} else if err == ErrInternal {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}
}


func authorization(token string) string {

	if err := validateToken(token); err != nil {
		return ErrNoSession
	}

	claims, er := ExtractClaims(token)
	if er != nil {
		return ErrInternal
	}
	username := claims.Username

	err := validateKey(username)
	if err == redis.Nil {
		return ErrNoSession
	} else if err != nil {
		return ErrInternal
	}

	return ""
}
