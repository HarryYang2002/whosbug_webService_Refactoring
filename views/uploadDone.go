package views

import (
	"errors"
	"fmt"
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
	fmt.Println(t)
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

	for i := 0; i < n; i++ {
		var nodes []NodesTable
		Db.Table("nodes").Find(&nodes)
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
		//var nodes1 []NodesTable
		var tnum int

		num, tnum := judge_object(temp2[i], nodes)
		fmt.Println("n:", tnum)
		if num != 0 { //有object
			if judge_change(temp3) == 1 { //没改
				t := nodes[tnum].OldConfidence
				nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
				nodes[tnum].NewConfidence = HightenConfidence(t)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
				//Db.M(&).Update("name", "hello")
			} else {
				t1 := nodes[tnum].OldConfidence
				nodes[tnum].OldConfidence = nodes[tnum].NewConfidence
				nodes[tnum].NewConfidence = CalculateConfidence(temp3, t1)
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("old_confidence", nodes[tnum].OldConfidence))
				fmt.Println(Db.Table("nodes").Where("table_id = ?", tnum).Update("new_confidence", nodes[tnum].NewConfidence))
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
			temp4.ObjectAdLine = temp2[i].AddedLine
			temp4.ObjectDeLine = temp2[i].DeletedLine
			temp4.ObjectNewLine = temp2[i].NewLine
			temp4.ObjectOldLine = temp2[i].OldLine
			fmt.Println(Db.Table("nodes").Create(&temp4).RowsAffected)

		}

	}
	context.Status(200)

}
