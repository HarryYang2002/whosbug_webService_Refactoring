package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
	. "webService_Refactoring/modules"
)

func CommitsDiffsCreate(context *gin.Context) {

	var t T4

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
	// 加入select子句, 现取消
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	commit := CommitsTable{}
	// 加入select子句, 现取消
	Db.Table("commits").First(&commit, "release_table_id = ?", temp1.TableId)
	n := len(t.UncountedObject)
	releaseId := temp1.TableId
	commitId := commit.TableId
	start := time.Now()
	// objectsSlice := make([]ObjectsTable, n) // 批量插入
	for i := 0; i < n; i++ {
		temp2 := ObjectsTable{}
		temp2.CommitTableId = int(commitId)
		temp2.ReleaseTableId = int(releaseId)
		temp2.FatherObjectId = t.UncountedObject[i].OldObjectId
		temp2.DeletedLine = t.UncountedObject[i].DeletedLineCount
		temp2.EndLine = t.UncountedObject[i].EndLine
		temp2.Hash = t.UncountedObject[i].Hash
		temp2.NewLine = t.UncountedObject[i].NewLineCount
		temp2.CurrentObjectId = t.UncountedObject[i].ObjectId
		temp2.ObjectPath = t.UncountedObject[i].Path
		temp2.OldLine = t.UncountedObject[i].OldLineCount
		temp2.Parameters = t.UncountedObject[i].Parameters
		temp2.StartLine = t.UncountedObject[i].StartLine
		temp2.AddedLine = t.UncountedObject[i].AddedLineCount
		// objectsSlice = append(objectsSlice, temp2) // 批量插入
		fmt.Println(Db.Table("objects").Create(&temp2).RowsAffected)
	}
	// Db.Table("objects").Create(&objectsSlice) // 批量插入
	context.Status(200)
	fmt.Printf("test time :%v", time.Since(start))

}
