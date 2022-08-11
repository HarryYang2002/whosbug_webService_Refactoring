package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"webService_Refactoring/modules"
)

func GetLastRelease(c *gin.Context) {
	var id modules.ProjectID
	if err := c.ShouldBind(&id); err != nil {
		err.Error()
	}
	dsn := "host=localhost user=postgres password=endata dbname=whosbugdemo port=60000 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	temp := modules.DbCreateProject{}
	res := db.Table("commits").Where("project_id = ?", id.Pid).First(&temp)
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
