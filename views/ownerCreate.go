package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"
)

func OwnerCreate(context *gin.Context) {
	var t GetConfidence
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	pid, err2 := strconv.Atoi(t.Project.Pid)
	if err2 != nil {
		context.Status(404)
		return
	}
	version := t.Release.Version
	temp := ProjectsTable{}
	res := db.Table("projects").First(&temp, "project_id = ?", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Project pid " + t.Project.Pid + " not exists",
		})
	}
	temp1 := ReleasesTable{}
	res1 := db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Release version " + t.Release.Version + " not exists",
		})
	}
	methods := t.Method
	n := len(methods)
	// TODO nodes表（原objects表）还未确定
	for i := 0; i < n; i++ {
		return
	}

}
