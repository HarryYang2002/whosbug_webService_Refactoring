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

func OwnerCreate(context *gin.Context) {
	var t GetConfidence
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	pid, err2 := strconv.Atoi(t.Project.Pid)
	if err2 != nil {
		context.Status(404)
		return
	}
	version := t.Release.Version
	temp := ProjectsTable{}
	res := Db.Table("projects").First(&temp, "project_id = ?", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Project pid " + t.Project.Pid + " not exists",
		})
		return
	}
	temp1 := ReleasesTable{}
	res1 := Db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Release version " + t.Release.Version + " not exists",
		})
		return
	}
	methods := t.Method
	n := len(methods)

	jsonResult := JsonRes{}
	var params []NodesTable
	for i := 0; i < n; i++ {
		methodId := methods[i].MethodId
		filePath := methods[i].Filepath
		parameters := methods[i].Parameters
		commitDemo := CommitsTable{}
		Db.Table("commits").First(&commitDemo, "release_version = ?", version)
		commitTableId := commitDemo.TableId
		var nodes []NodesTable
		//数据库中查找所有符合条件的数据
		Db.Table("nodes").Find(&nodes, "commit_table_id in (?)", commitTableId)
		if len(nodes) == 0 {
			context.Status(400)
			return
		}
		//第一次筛选
		var methods2 []NodesTable
		for x := 0; x < len(nodes); x++ {
			if nodes[i].CurrentObjectId == methodId {
				methods2 = append(methods2, nodes[i])
			}
		}
		if len(methods2) == 0 {
			fmt.Println("Get objects error:")
			jsonResult.Message = "No such objects in release: version: " + version
			jsonResult.Status = "may be ok"
			fmt.Println(jsonResult)
			continue
		}
		//第二次筛选
		var path []NodesTable
		for x := 0; x < len(nodes); x++ {
			if nodes[i].ObjectPath == filePath {
				path = append(path, nodes[i])
			}
		}
		if len(path) == 0 {
			fmt.Println("Get objects error:")
			jsonResult.Message = "No such objects in path: filepath: " + filePath + " here's results with id"
			jsonResult.Object = nodes
			jsonResult.Status = "may be ok"
			fmt.Println(jsonResult)
			continue
		}
		//第三次筛选
		for x := 0; x < len(path); x++ {
			if path[i].ObjectParameters == parameters {
				params = append(params, path[i])
			}
		}
		if len(params) == 0 {
			fmt.Println("Get objects error:")
			jsonResult.Message = "No such objects in params: " + parameters + " here's results with path"
			jsonResult.Object = path
			jsonResult.Status = "may be ok"
			fmt.Println(jsonResult)
			continue
		}
	}
	var objectInfos []ObjectInfo
	for i := 0; i < len(params); i++ {
		objectInfo := ObjectInfo{}
		objects := ObjectsTable{}
		Db.Table("objects").First(&objects, "table_id = ?", params[i].ObjectTableId)
		objectInfo.hash = objects.Hash
		objectInfo.objectId = objects.CurrentObjectId
		objectInfo.oldObjectId = objects.FatherObjectId
		objectInfo.confidence = params[i].NewConfidence
		objectInfo.parameters = objects.Parameters
		objectInfo.startLine = objects.StartLine
		objectInfo.endLine = objects.EndLine
		objectInfo.oldlineCount = objects.OldLine
		objectInfo.newlineCount = objects.NewLine
		objectInfo.deletedlineCount = objects.DeletedLine
		objectInfo.addedLineCount = objects.AddedLine
		objectInfos = append(objectInfos, objectInfo)
	}

	OriginInfo := GetBugOrigin(objectInfos)
	context.JSON(http.StatusOK, gin.H{
		"ownerInfo": OriginInfo,
	})
}
