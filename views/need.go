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

type ObjectInfo struct {
	objectId         string  `json:"object_id"`         //Object的函数唯一标识符
	oldObjectId      string  `json:"old_object_id"`     //Object的父类唯一表示符
	confidence       float64 `json:"confidence"`        //置信度
	parameters       string  `json:"parameters"`        //方法的参数特征
	oldlineCount     int     `json:"oldline_count"`     //旧行数
	newlineCount     int     `json:"newline_count"`     //新行数
	deletedlineCount int     `json:"deletedline_count"` //移除行数
	addedLineCount   int     `json:"added_line_count"`  //新增行数
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
	object    ObjectInfo         `json:"object"`
	wrongRate float64            `json:"wrong_rate"`
	owners    map[string]float64 `json:"owners"`
}

type TreeNode struct {
	sonNum int
	object ObjectInfo
	childs []TreeNode
}

//  @param objectId
//  @return []historyInfo
//  返回的切片要按时间顺序排，最新的commit及其对应object放在索引0
func getHistory(objectId string) (result []HistoryInfo) {
	var temp []NodesTable
	res2 := Db.Table("nodes").Where("current_object_id = ? ", objectId).Find(&temp)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		return
	}
	n := len(temp)
	for i := 0; i < n; i++ {
		var temp1 CommitsTable
		tableId := temp[i].CommitTableID
		res := Db.Table("commits").Where("table_id = ? ", tableId).First(&temp1)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return
		}
		var temp3 ObjectHistoryInfo
		var temp4 CommitInfo
		temp3.addedLineCount = temp[i].ObjectAdLine
		temp3.deletedlineCount = temp[i].ObjectDeLine
		temp3.newlineCount = temp[i].ObjectNewLine
		temp3.oldlineCount = temp[i].ObjectOldLine
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
func getChain(objectID string) (node TreeNode) {
	temp := ObjectsTable{}
	Db.Table("objects").First(&temp, "current_object_id = ?", objectID)
	node.object = ObjectInfo{temp.CurrentObjectID, temp.FatherObjectID, 0,
		temp.Parameters, temp.OldLine, temp.NewLine,
		temp.DeletedLine, temp.AddedLine}
	var tempChilds []ObjectsTable
	Db.Table("objects").Find(&tempChilds, "father_object_id in (?)", objectID)
	for i := range tempChilds {
		node.childs = append(node.childs, getChain(tempChilds[i].CurrentObjectID))
	}
	return
}
