package whosbug

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// OwnerCreate
// @param context *gin.Context
// @Description 根据第三方传递的报错信息找出责任人
// @author: HarryYang 2022-08-23 14:45:16
func OwnerCreate(context *gin.Context) {
	var t GetConfidence
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	pid := t.Project.Pid
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
	res1 := Db.Table("releases").Select("table_id").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Release version " + t.Release.Version + " not exists",
		})
		return
	}
	releaseTableId := temp1.TableID
	methods := t.Method
	n := len(methods)

	jsonResult := JsonRes{}
	var params []NodesTable
	for i := 0; i < n; i++ {
		methodId := methods[i].MethodId
		filePath := methods[i].Filepath
		parameters := methods[i].Parameters
		commitDemo := CommitsTable{}
		Db.Table("commits").Select("table_id").First(&commitDemo, "release_table_id = ?", releaseTableId)
		commitTableId := commitDemo.TableID
		var nodes []NodesTable
		//数据库中查找所有符合条件的数据
		Db.Table("nodes").Find(&nodes, "commit_table_id in (?)", commitTableId)
		if len(nodes) == 0 {
			context.Status(400)
			continue
		}
		//第一次筛选
		var methods2 []NodesTable
		for x := 0; x < len(nodes); x++ {
			if nodes[x].CurrentObjectID == methodId {
				methods2 = append(methods2, nodes[x])
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
		for x := 0; x < len(methods2); x++ {
			if methods2[x].ObjectPath == filePath {
				path = append(path, methods2[x])
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
			if path[x].ObjectParameters == parameters {
				params = append(params, path[x])
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
		objectInfo.objectId = params[i].CurrentObjectID
		objectInfo.oldObjectId = params[i].FatherObjectID
		objectInfo.confidence = params[i].NewConfidence
		objectInfo.parameters = params[i].ObjectParameters
		objectInfo.newlineCount = params[i].ObjectNewLine
		objectInfo.oldlineCount = params[i].ObjectOldLine
		objectInfo.deletedlineCount = params[i].ObjectDeLine
		objectInfo.addedLineCount = params[i].ObjectAdLine
		objectInfos = append(objectInfos, objectInfo)
	}

	OriginInfo := GetBugOrigin(objectInfos)
	fmt.Println(OriginInfo)
	var knum []string
	var vnum []float64
	for k, v := range OriginInfo {
		knum = append(knum, k)
		vnum = append(vnum, v)
	}
	type testss struct {
		Name   string  `json:"name" binding:"required"`
		Weight float64 `json:"weight" binding:"required"`
	}
	var tt []testss
	var tt1 testss
	for i := 0; i < len(OriginInfo); i++ {
		tt1.Name = knum[i]
		tt1.Weight = vnum[i]
		tt = append(tt, tt1)
	}
	context.JSON(http.StatusOK, gin.H{
		"owner": tt,
	})
}
