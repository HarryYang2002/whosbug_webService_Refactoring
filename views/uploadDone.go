package views

import (
	"errors"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommitsUploadDoneCreate(context *gin.Context) {

	var t T

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

	temp2 := ObjectsTable{}
	Db.Table("objects").Model(&ObjectsTable{}).Find(&temp2)

	var nodes []NodesTable
	Db.Table("nodes").Model(&NodesTable{}).Find(&nodes)

	temp3 := UncalculateObjectInfo{}

	temp3.addedLineCount = temp2.AddedLine
	temp3.deletedlineCount = temp2.DeletedLine
	temp3.endLine = temp2.EndLine
	temp3.hash = temp2.Hash
	temp3.newlineCount = temp2.NewLine
	temp3.objectId = temp2.CurrentObjectId
	temp3.oldObjectId = temp2.FatherObjectId
	temp3.oldlineCount = temp2.OldLine
	temp3.parameters = temp2.Parameters
	temp3.startLine = temp2.StartLine
	var nodes1 []NodesTable
	var tnum int
	num, nodes1, tnum := judge_object(temp2, nodes)
	if num != 0 { //有object
		if judge_change(temp3) == 1 { //没改
			nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
			nodes[tnum].NewConfidence = HightenConfidence(nodes1[num].NewConfidence)
		} else {
			nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
			nodes[tnum].NewConfidence = CalculateConfidence(temp3, nodes1[num].NewConfidence)
		}
	} else {
		temp4 := NodesTable{}
		temp4.CommitTableId = temp2.CommitTableId
		temp4.CurrentObjectId = temp2.CurrentObjectId
		temp4.FatherObjectId = temp2.FatherObjectId
		temp4.NewConfidence = CalculateConfidence(temp3, 0)
		temp4.ObjectPath = temp2.ObjectPath
		temp4.ObjectTableId = int(temp2.TableId) //?
		temp4.ObjectsParameters = temp2.Parameters
		temp4.OldConfidence = 0
		Db.Table("nodes").Create(&temp4)

	}

	// temp3 := ObjectsTable{}
	// lastCommitHash := t.Release.CommitHash
	// errs := db.Table("objects").First(&temp3, "release_version = ? and hash = ?", version, lastCommitHash) //数据传到temp3里面，拿temp3去做算法
	// if errs != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Delete error",
	// 	})
	// 	return
	// }
	context.Status(200)

}
