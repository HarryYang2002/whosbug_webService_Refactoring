package views

import (
	"errors"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommitsUploadDoneCreate(context *gin.Context) {

	var t T

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid, err := strconv.Atoi(t.Project.Pid)
	if err != nil {
		context.Status(404)
	}
	version := t.Release.Version
	temp := ProjectsTable{}
	res := db.Table("projects").First(&temp, "project_id = ? ", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp1 := ReleasesTable{}
	res1 := db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp3 := ObjectsTable{}
	lastCommitHash := t.Release.CommitHash
	errs := db.Table("objects").First(&temp3, "release_version = ? and hash = ?", version, lastCommitHash) //数据传到temp3里面，拿temp3去做算法
	if errs != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
		})
		return
	}
	context.Status(200)

}
