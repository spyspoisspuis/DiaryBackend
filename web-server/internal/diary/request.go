package diary

import (
	"database/sql"
	"net/http"
	"web-server/internal/authen"
	"web-server/internal/db"
	"web-server/internal/util"

	"github.com/gin-gonic/gin"
)

func AddDiary(c *gin.Context) {
	var inp util.DiaryStruct
	if err := c.ShouldBind(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}

	err := db.AddDiary(inp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)

}

func UpdateDiary(c *gin.Context) {
	var inp util.DiaryStruct
	if err := c.ShouldBind(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}

	err := db.DeleteDiary(inp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	err = db.AddDiary(inp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func GetDiary(c *gin.Context) {
	var inp GetDiaryInput
	if err := c.ShouldBind(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrBadRequest})
		return
	}

	diaryStruct, err := db.GetDiary(*inp.Writer, *inp.Week)
	if err == sql.ErrNoRows {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username, err := authen.RetreiveUsernameFromHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ErrInternal})
		return
	}

	// not the same person check if status is confirm
	if diaryStruct.Status != "confirm" {
		if diaryStruct.Writer != username {
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Diary": diaryStruct})
}