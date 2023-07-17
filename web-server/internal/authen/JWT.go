package authen

import (
	"fmt"
	"web-server/internal/dstore"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Validate the token 
func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		signed := viper.GetString("jwt.signed")
		return []byte(signed), nil
	})

	return err
}


// Extract the claims from token in this point we have only username
func ExtractClaims(tokenString string) (*UserClaims, error) {
	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method and return the key to verify the signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		signed := viper.GetString("jwt.signed")
		return []byte(signed), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		// handle token validation error
		return nil, err
	}
	return claims, nil
}

// Validate the key of token aka username whether the token expire ? 
func validateKey(username string) error {
	_, err := dstore.GetToken(username)
	return err
}

func GenerateJWTtoken(username string) (string, error) {
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed := viper.GetString("jwt.signed")

	ss, err := token.SignedString([]byte(signed))
	if err != nil {
		return "", err
	}
	if err = dstore.LoginSession(username, ss, 24*time.Hour); err != nil {
		return "", err
	}
	return ss, nil
}

// Retrieve username from token inside header 
func RetreiveUsernameFromHeader(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := ExtractClaims(token)
	if err != nil {
		return "", err
	}
	username := claims.Username
	return username, nil
}
