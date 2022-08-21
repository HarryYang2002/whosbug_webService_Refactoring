package whosbug

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

func GetLastRelease(c *gin.Context) {
	var id ProjectID
	if err := c.ShouldBind(&id); err != nil {
		err.Error()
	}
	projectid := ProjectsTable{}
	Db.Table("projects").Select("table_id").Where("project_id = ?", id.Pid).First(&projectid)
	projectTableId := projectid.TableID
	temp := ReleasesTable{}
	res := Db.Table("releases").Select("release_version").Where("project_table_id = ?", projectTableId).First(&temp)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		info := "The project with pid = " + id.Pid + " does not exists."
		c.JSON(http.StatusNotFound, gin.H{
			"error": info,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"lastReleaseVersion": temp.ReleaseVersion,
	})

}
