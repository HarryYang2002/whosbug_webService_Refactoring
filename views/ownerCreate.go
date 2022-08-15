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
	res := db.Table("projects").First(&temp, "project_id = ?", pid)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Project pid " + t.Project.Pid + " not exists",
		})
	}
	temp1 := ReleasesTable{}
	res1 := db.Table("releases").First(&temp1, "release_version = ?", version)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Release version " + t.Release.Version + " not exists",
		})
	}
	methods := t.Method
	n := len(methods)

	jsonres := JsonRes{}
	for i := 0; i < n; i++ {
		methodId := methods[i].MethodId
		filePath := methods[i].Filepath
		parameters := methods[i].Parameters
		commitDemo := CommitsTable{}
		db.Table("commits").First(&commitDemo, "release_version = ?", version)
		commitTableId := commitDemo.TableId
		var nodes []NodesTable
		//数据库中查找所有符合条件的数据
		db.Table("objects").Model(&NodesTable{}).Where("commit_table_id in (?)", commitTableId).Find(&nodes)
		if len(nodes) == 0 {
			context.Status(400)
			return
		}
		//第一次筛选
		var methods []NodesTable
		for x := 0; x < len(nodes); x++ {
			if nodes[i].CurrentObjectId == methodId {
				methods = append(methods, nodes[i])
			}
		}
		if len(methods) == 0 {
			fmt.Println("Get objects error:")
			jsonres.Message = "No such objects in release: version: " + version
			jsonres.Status = "may be ok"
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
			jsonres.Message = "No such objects in path: filepath: " + filePath + " here's results with id"
			jsonres.Object = nodes
			jsonres.Status = "may be ok"
			continue
		}
		//第三次筛选
		var params []NodesTable
		for x := 0; x < len(path); x++ {
			if path[i].ObjectsParameters == parameters {
				params = append(params, path[i])
			}
		}
		if len(params) == 0 {
			fmt.Println("Get objects error:")
			jsonres.Message = "No such objects in params: " + parameters + " here's results with path"
			jsonres.Object = path
			jsonres.Status = "may be ok"
			continue
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"res": jsonres,
	})

}
