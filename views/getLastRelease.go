package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"
)

func GetLastRelease(c *gin.Context) {
	var id ProjectId
	if err := c.ShouldBind(&id); err != nil {
		err.Error()
	}
	temp := ReleasesTable{}
	res := Db.Table("releases").Where("project_id = ?", id.Pid).First(&temp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		info := "The project with pid = " + strconv.Itoa(id.Pid) + " does not exists."
		c.JSON(http.StatusNotFound, gin.H{
			"error": info,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"lastReleaseVersion": temp.ReleaseVersion,
	})

}
