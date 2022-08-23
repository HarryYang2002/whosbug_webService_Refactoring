package whosbug

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CommitsDiffsCreate
// @param context *gin.Context
// @Description 上传object信息
// @author: TongLei 2022-08-23 14:41:09
func CommitsDiffsCreate(context *gin.Context) {

	var t T4

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//获取pid，version
	pid := t.Project.Pid
	version := t.Release.Version
	temp := ProjectsTable{}
	//从数据库中获取数据
	res := Db.Table("projects").First(&temp, "project_id = ? ", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	temp1 := ReleasesTable{}
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	//上传objects数据
	commit := CommitsTable{}
	Db.Table("commits").First(&commit, "release_table_id = ?", temp1.TableID)
	n := len(t.UncountedObject)
	releaseId := temp1.TableID
	commitId := commit.TableID
	for i := 0; i < n; i++ {
		temp2 := ObjectsTable{}
		temp2.CommitTableID = int(commitId)
		temp2.ReleaseTableID = int(releaseId)
		temp2.FatherObjectID = t.UncountedObject[i].OldObjectId
		temp2.DeletedLine = t.UncountedObject[i].DeletedLineCount
		temp2.EndLine = t.UncountedObject[i].EndLine
		temp2.Hash = t.UncountedObject[i].Hash
		temp2.NewLine = t.UncountedObject[i].NewLineCount
		temp2.CurrentObjectID = t.UncountedObject[i].ObjectId
		temp2.ObjectPath = t.UncountedObject[i].Path
		temp2.OldLine = t.UncountedObject[i].OldLineCount
		temp2.Parameters = t.UncountedObject[i].Parameters
		temp2.StartLine = t.UncountedObject[i].StartLine
		temp2.AddedLine = t.UncountedObject[i].AddedLineCount
		fmt.Println(Db.Table("objects").Create(&temp2).RowsAffected)
	}

	context.Status(200)

}
