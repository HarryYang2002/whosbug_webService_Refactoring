package views

import (
	"errors"
	"gorm.io/gorm"
	. "webService_Refactoring/modules"
)

type CommitInfo struct {
	commitHash   string
	commitAuthor string
	commitEmail  string
	commitTime   string
}

type ObjectInfo struct { //objects
	hash             string  //Object所属的Commit
	objectId         string  //Object的函数唯一标识符  现版本
	oldObjectId      string  //Object的父类唯一表示符  老
	confidence       float64 //置信度   nodes表
	parameters       string  //方法的参数特征
	startLine        int     //起始行
	endLine          int     //结束行
	oldlineCount     int     //旧行数
	newlineCount     int     //新行数
	deletedlineCount int     //移除行数
	addedLineCount   int     //新增行数
}

type UncalculateObjectInfo struct {
	hash             string //Object所属的Commit
	objectId         string //Object的函数唯一标识符
	oldObjectId      string //Object的父类唯一表示符
	parameters       string //方法的参数特征
	startLine        int    //起始行
	endLine          int    //结束行
	oldlineCount     int    //旧行数
	newlineCount     int    //新行数
	deletedlineCount int    //移除行数
	addedLineCount   int    //新增行数
}

type OwnerInfo struct {
	author string  //责任人名称
	email  string  //邮箱
	weight float64 //权重
}

type ObjectHistoryInfo struct {
	oldlineCount     int //旧行数
	newlineCount     int //新行数
	deletedlineCount int //移除行数
	addedLineCount   int //新增行数
}

type HistoryInfo struct {
	commitHistory CommitInfo
	objectHistory ObjectHistoryInfo
}

type bugOriginInfo struct {
	object    ObjectInfo
	wrongRate float64
	owners    map[string]float64
}

type TreeNode struct {
	object ObjectInfo
	childs []ObjectInfo
}

//  @param objectId
//  @return []historyInfo
//  返回的切片要按时间顺序排，最新的commit及其对应object放在索引0
func getHistory(objectId string) (result []HistoryInfo) {
	var temp []ObjectsTable
	res2 := Db.Table("objects").Where("current_object_id = ? ", objectId).Find(&temp)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		return
	}
	n := len(temp)
	for i := 0; i < n; i++ {
		var temp1 CommitsTable
		tableId := temp[i].CommitTableId
		res := Db.Table("commits").Where("table_id = ? ", tableId).First(&temp1)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return
		}
		var temp3 ObjectHistoryInfo
		var temp4 CommitInfo
		temp3.addedLineCount = temp[i].AddedLine
		temp3.deletedlineCount = temp[i].DeletedLine
		temp3.newlineCount = temp[i].NewLine
		temp3.oldlineCount = temp[i].OldLine
		temp4.commitAuthor = temp1.Author
		temp4.commitEmail = temp1.Email
		temp4.commitHash = temp1.Hash
		temp4.commitTime = temp1.Time
		var temp5 HistoryInfo
		temp5.objectHistory = temp3
		temp5.commitHistory = temp4
		result = append(result, temp5)
	}

	return result

}

//  @param objectId 函数的id
//  @return	chainNode 该函数所在的定义链的根结点
func getChain(objectId string) (node TreeNode) {
	//var chlids []ObjectInfo
	//nodeModule := NodesTable{}
	//Db.Table("nodes").First(&nodeModule, "current_object_id = ?", objectId)
	//if nodeModule.FatherObjectId != "" {
	//	child := ObjectInfo{}
	//	child.hash =
	//	id := nodeModule.FatherObjectId
	//	Db.Table("nodes").First(&nodeModule, "current_object_id = ?", id)
	//}
	return
}
