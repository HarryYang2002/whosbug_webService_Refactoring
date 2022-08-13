package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"
)

// CreateProjectRelease 生成project&release
func CreateProjectRelease(context *gin.Context) {
	var t T
	err1 := context.ShouldBind(&t)
	if err1 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return
	}
	pid, err := strconv.Atoi(t.Project.Pid)
	if err != nil {
		context.Status(404)
	}
	releaseVersion := t.Release.Version
	releaseHash := t.Release.CommitHash
	// 数据库查询pid，若存在且数据库中last_commit_hash 为传递的last_commit_hash
	// 不新建project并返回404
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		err2.Error()
	}
	project := ProjectsTable{}
	res := db.Table("projects").Where("project_id = ?", pid).First(&project)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		project.ProjectId = pid
		fmt.Println(db.Table("projects").Create(&project).RowsAffected)
	}
	release := ReleasesTable{}
	res2 := db.Table("releases").Where("release_version = ? "+
		"and last_commit_hash = ?", releaseVersion, releaseHash).First(&release)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		release.ProjectId = pid
		release.ReleaseVersion = releaseVersion
		release.LastCommitHash = releaseHash
		fmt.Println(db.Table("releases").Create(&release).RowsAffected)
	} else {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "The Project and Release already exist, update the commit pid " + t.Project.Pid +
				" release: " + t.Release.Version + ", commit_hash: " + t.Release.CommitHash,
		})
		return
	}
	context.Status(201)
}
