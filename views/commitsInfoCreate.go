package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"
)

// 写入数据异常

func CommitsInfoCreate(context *gin.Context) {

	var t T2

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid, err2 := strconv.Atoi(t.Project.Pid)
	if err2 != nil {
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
	releaseTableId := temp1.TableId
	n := len(t.Commit)
	for i := 0; i < n; i++ {
		temp2 := CommitsTable{}
		temp2.ReleaseTableId = int(releaseTableId)
		temp2.Hash = t.Commit[i].Hash
		temp2.Author = t.Commit[i].Author
		temp2.Email = t.Commit[i].Email
		temp2.Time = t.Commit[i].Email
		fmt.Println(db.Table("commits").Create(&temp2).RowsAffected)
	}
	context.Status(200)

}
