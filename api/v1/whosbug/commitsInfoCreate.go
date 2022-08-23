package whosbug

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CommitsInfoCreate
// @param context *gin.Context
// @Description 上传commit信息
// @author: TongLei 2022-08-23 14:28:11
func CommitsInfoCreate(context *gin.Context) {

	var t T2

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid := t.Project.Pid
	version := t.Release.Version
	temp := ProjectsTable{}
	res := Db.Table("projects").First(&temp, "project_id = ? ", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp1 := ReleasesTable{}
	res1 := Db.Table("releases").Select("table_id").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	//上传commits数据
	releaseTableId := temp1.TableID
	n := len(t.Commit)
	commitsSlice := make([]CommitsTable, n) // 批量插入
	for i := 0; i < n; i++ {
		temp2 := CommitsTable{}
		temp2.ReleaseTableID = int(releaseTableId)
		temp2.Hash = t.Commit[i].Hash
		temp2.Author = t.Commit[i].Author
		temp2.Email = t.Commit[i].Email
		temp2.Time = t.Commit[i].Time
		fmt.Println(Db.Table("commits").Create(&temp2).RowsAffected)
		temp2.Time = t.Commit[i].Time
		commitsSlice[i] = temp2
	}
	Db.Table("commits").Create(&commitsSlice)
	context.Status(200)

}
