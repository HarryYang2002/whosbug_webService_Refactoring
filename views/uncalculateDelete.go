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
	pid, err2 := strconv.Atoi(t.Project.Pid)
	if err2 != nil {
		context.Status(404)
	}
	version := t.Release.Version
	//连接数据库
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		err2.Error()
	}
	//以pid去找
	project := ProjectsTable{}
	res := db.Table("projects").Where("project_id = ?", pid).First(&project)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Project get fails",
			"detail": "no such project:" + strconv.Itoa(pid),
		})
		return
	}
	//以version去找
	release := ReleasesTable{}
	res2 := db.Table("releases").Where("release_version = ? and project_id = ?", version, pid).First(&release)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Release get fails",
			"detail": "no such release:" + version,
		})
		return
	}
	realRelease := ReleasesTable{}
	uncounted := UncountedObjectsTable{}
	commit := CommitsTable{}
	db.Table("releases").First(&realRelease, "release_version = ?", version)
	releaseId := realRelease.TableId
	db.Table("commits").First(&commit, "release_table_id = ?", releaseId)
	uncountedId := commit.TableId
	res5 := db.Table("uncounted_objects").Delete(&uncounted, "commit_table_id = ?", uncountedId)
	if res5.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"Success": "Success",
	})
}
