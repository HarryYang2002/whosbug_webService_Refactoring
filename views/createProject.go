package views

import (
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
	// 数据库查询pid，若存在且数据库中last_commit_hash 为传递的last_commit_hash
	// 不新建project并返回404
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		err2.Error()
	}
	temp := DbCreateProject{}
	//db.Where("project_id = ?", pid).First(&temp)
	db.Table("commits").Where("project_id = ? and release_version = ?", pid, t.Release.Version).Find(&temp)
	if temp.ReleaseVersion != "" && temp.ReleaseVersion == t.Release.Version {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "The Project and Release already exist, update the commit pid " + t.Project.Pid +
				" release: " + t.Release.Version + ", commit_hash: " + t.Release.CommitHash,
		})
		return
	}
	release := t.Release.Version
	commitHash := t.Release.CommitHash
	temp.ProjectId = pid
	temp.ReleaseVersion = release
	temp.PvLastCommitHash = commitHash
	//新建project并存储进数据库中
	fmt.Println(db.Table("commits").Create(&temp).RowsAffected)
	context.Status(201)
}
