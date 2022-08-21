package whosbug

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// CommitsDiffsCreate 在数据库中创建commitdiff
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
	res1 := Db.Table("releases").Select("table_id").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	commit := CommitsTable{}
	Db.Table("commits").Select("table_id").First(&commit, "release_table_id = ?", temp1.TableID)
	n := len(t.UncountedObject)
	releaseId := temp1.TableID
	commitId := commit.TableID
	objectsSlice := make([]ObjectsTable, n) // 批量插入
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
		objectsSlice[i] = temp2
	}
	Db.Table("objects").Create(&objectsSlice)
	context.Status(200)

}
