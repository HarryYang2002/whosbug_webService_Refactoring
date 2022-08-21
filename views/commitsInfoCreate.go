package views

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	pid := t.Project.Pid
	version := t.Release.Version
	temp := ProjectsTable{}
	res := Db.Table("projects").First(&temp, "project_id = ? ", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp1 := ReleasesTable{}
	// 插入select子句, 现取消
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	releaseTableId := temp1.TableId
	n := len(t.Commit)
	CommitsSlice := make([]CommitsTable, n) // 批量
	for i := 0; i < n; i++ {
		temp2 := CommitsTable{}
		temp2.ReleaseTableId = int(releaseTableId)
		temp2.Hash = t.Commit[i].Hash
		temp2.Author = t.Commit[i].Author
		temp2.Email = t.Commit[i].Email
		temp2.Time = t.Commit[i].Email
		// fmt.Println(Db.Table("commits").Create(&temp2).RowsAffected)
		CommitsSlice = append(CommitsSlice, temp2) // 批量
	}
	Db.Table("commits").Create(&CommitsSlice) // 批量
	context.Status(200)

}
