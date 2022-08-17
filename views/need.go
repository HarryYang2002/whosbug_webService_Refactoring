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
	//hash             string  //Object所属的Commit
	objectId    string  //Object的函数唯一标识符
	oldObjectId string  //Object的父类唯一表示符
	confidence  float64 //置信度
	parameters  string  //方法的参数特征
	//startLine        int     //起始行
	//endLine          int     //结束行
	oldlineCount     int //旧行数
	newlineCount     int //新行数
	deletedlineCount int //移除行数
	addedLineCount   int //新增行数
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

//TODO 数据库中查询不到old_object_id，已经确定为ci插件传递错误

//  @param objectId 函数的id
//  @return	chainNode 该函数所在的定义链的根结点
func getChain(objectId string) (node TreeNode) {
	var childs []ObjectInfo
	var father ObjectInfo
	//循环，结束条件为该条数据没有FatherObjectId
	//for {
	//	nodeModule := NodesTable{}
	//	res := Db.Table("nodes").First(&nodeModule, "current_object_id = ?", objectId)
	//	if errors.Is(res.Error,gorm.ErrRecordNotFound) {
	//
	//	}
	//	//存在FatherObjectId，将其加入childs切片
	//	if nodeModule.FatherObjectId != "" {
	//		child := ObjectInfo{}
	//		child.addedLineCount = nodeModule.ObjectAdLine
	//		child.objectId = nodeModule.CurrentObjectId
	//		child.newlineCount = nodeModule.ObjectNewLine
	//		child.oldlineCount = nodeModule.ObjectOldLine
	//		child.deletedlineCount = nodeModule.ObjectDeLine
	//		child.oldObjectId = nodeModule.FatherObjectId
	//		child.parameters = nodeModule.ObjectParameters
	//		child.confidence = nodeModule.NewConfidence
	//		childs = append(childs, child)
	//		objectId = nodeModule.FatherObjectId
	//	} else {
	//		father.addedLineCount = nodeModule.ObjectAdLine
	//		father.deletedlineCount = nodeModule.ObjectDeLine
	//		father.confidence = nodeModule.NewConfidence
	//		father.objectId = nodeModule.CurrentObjectId
	//		father.objectId = nodeModule.FatherObjectId
	//		father.parameters = nodeModule.ObjectParameters
	//		father.newlineCount = nodeModule.ObjectNewLine
	//		father.oldlineCount = nodeModule.ObjectOldLine
	//		break
	//	}
	//}
	node.childs = childs
	node.object = father
	return node
}
