package authen
import (
	"github.com/dgrijalva/jwt-go"
)
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
    jwt.StandardClaims
    Username string `json:"username"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}