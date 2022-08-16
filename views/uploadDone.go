package views

import (
	"errors"
	"net/http"
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
	pid := t.Project.Pid
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

	//temp2 := ObjectsTable{}
	var temp2 []ObjectsTable
	Db.Table("objects").Find(&temp2)
	n := len(temp2)
	var nodes []NodesTable
	Db.Table("nodes").Find(&nodes)

	for i := 0; i < n; i++ {
		temp3 := UncalculateObjectInfo{}
		temp3.addedLineCount = temp2[i].AddedLine
		temp3.deletedlineCount = temp2[i].DeletedLine
		temp3.endLine = temp2[i].EndLine
		temp3.hash = temp2[i].Hash
		temp3.newlineCount = temp2[i].NewLine
		temp3.objectId = temp2[i].CurrentObjectId
		temp3.oldObjectId = temp2[i].FatherObjectId
		temp3.oldlineCount = temp2[i].OldLine
		temp3.parameters = temp2[i].Parameters
		temp3.startLine = temp2[i].StartLine
		var nodes1 []NodesTable
		var tnum int

		num, nodes1, tnum := judge_object(temp2[i], nodes)
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
			temp4.CommitTableId = temp2[i].CommitTableId
			temp4.CurrentObjectId = temp2[i].CurrentObjectId
			temp4.FatherObjectId = temp2[i].FatherObjectId
			temp4.NewConfidence = CalculateConfidence(temp3, 0)
			temp4.ObjectPath = temp2[i].ObjectPath
			temp4.ObjectTableId = int(temp2[i].TableId) //?
			temp4.ObjectParameters = temp2[i].Parameters
			temp4.OldConfidence = 0
			Db.Table("nodes").Create(&temp4)

		}

	}
	context.Status(200)

}
