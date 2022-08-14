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

//TODO 逻辑貌似有问题

func CommitsDiffsCreate(context *gin.Context) {

	var t T4

	err := context.ShouldBind(&t)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	pid, err := strconv.Atoi(t.Project.Pid)
	if err != nil {
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
	res1 := db.Table("release").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.Status(400)
		return
	}
	realRelease := ReleasesTable{}
	commit := CommitsTable{}
	temp2 := UncountedObjectsTable{}
	n := len(t.UncountedObject)
	for i := 0; i < n; i++ {
		db.Table("releases").First(&realRelease, "release_version = ?", version)
		releaseId := realRelease.TableId
		db.Table("commits").First(&commit, "release_table_id = ?", releaseId)
		commitId := commit.TableId
		temp2.CommitTableId = int(commitId)
		temp2.ReleaseTableId = int(releaseId)
		temp2.OldObjectId = t.UncountedObject[i].OldObjectId
		temp2.DeleteLine = t.UncountedObject[i].DeletedLineCount
		temp2.EndLine = t.UncountedObject[i].EndLine
		temp2.Hash = t.UncountedObject[i].Hash
		temp2.NewLine = t.UncountedObject[i].NewLineCount
		temp2.NewObjectId = t.UncountedObject[i].ObjectId
		temp2.ObjectPath = t.UncountedObject[i].Path
		temp2.OldLine = t.UncountedObject[i].OldLineCount
		temp2.Parameters = t.UncountedObject[i].Parameters
		temp2.StartLine = t.UncountedObject[i].StartLine

		fmt.Println(db.Table("uncounted_objects").Create(&temp2).RowsAffected)
	}

	context.Status(200)

}
