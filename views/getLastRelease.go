package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
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
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	temp := ReleasesTable{}
	res := db.Table("releases").Where("project_id = ?", id.Pid).First(&temp)
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
