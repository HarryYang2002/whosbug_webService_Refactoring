package whosbug

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// UncalculateDelete 接收完数据之后删除为计算的object信息
func UncalculateDelete(context *gin.Context) {
	//接收数据
	var t T
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//提取pid、version
	pid := t.Project.Pid
	version := t.Release.Version
	//以pid去找
	project := ProjectsTable{}
	res := Db.Table("projects").Where("project_id = ?", pid).First(&project)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Project get fails",
			"detail": "no such project:" + pid,
		})
		return
	}
	//以version去找
	release := ReleasesTable{}
	res2 := Db.Table("releases").Where("release_version = ? and project_table_id = ?", version, project.TableID).First(&release)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Release get fails",
			"detail": "no such release:" + version,
		})
		return
	}
	realRelease := ReleasesTable{}
	uncounted := ObjectsTable{}
	commit := CommitsTable{}
	Db.Table("releases").Select("table_id").First(&realRelease, "release_version = ?", version)
	releaseId := realRelease.TableID
	Db.Table("commits").Select("table_id").First(&commit, "release_table_id = ?", releaseId)
	uncountedId := commit.TableID
	res5 := Db.Table("objects").Delete(&uncounted, "commit_table_id = ?", uncountedId)
	if res5.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
			"err":   res5.Error.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"Success": "Success",
	})
}
