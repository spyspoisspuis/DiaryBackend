package authen

import (
	"database/sql"
	"net/http"
	"web-server/internal/db"
	"web-server/internal/dstore"
	"web-server/internal/util"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Logout(c *gin.Context) {
	username, er := RetreiveUsernameFromHeader(c)
	if er != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
		return
	}

	if err := dstore.RemoveToken(username); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func Login(c *gin.Context) {
	var creds LoginInput
	if err := c.ShouldBind(&creds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}

	dbName, err := db.GetUsername(creds.Username)
	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": ErrIncorrectUsername})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}

	encrpyted_Password, err := db.GetPasswordFromUsername(dbName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}

	if er := bcrypt.CompareHashAndPassword([]byte(encrpyted_Password), []byte(creds.Password)); er != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": ErrIncorrectPassword})
		return
	}

	var ss string
	err = validateKey(dbName)
	if err == redis.Nil {
		s, e := GenerateJWTtoken(dbName)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			return
		}
		ss = s
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	} else {
		s, e := dstore.GetToken(dbName)
		if e != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			return
		}
		ss = s
	}

	c.JSON(http.StatusOK, gin.H{"message": MessageSuccessLogin, "token": ss, "username": dbName})

}

func GetUsername(c *gin.Context) {
	username, err := RetreiveUsernameFromHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func Register(c *gin.Context) {
	var creds RegisterInput
	if err := c.ShouldBind(&creds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}

	dbName, err := db.GetUsername(creds.Username)
	if err != nil && err != sql.ErrNoRows {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}
	if dbName != "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": ErrDuplicateUsername})
		return
	}

	encryptPass, err := util.HashPassword(creds.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}

	err = db.InsertUser(creds.Username, encryptPass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": MessageSuccessRegister})

}
