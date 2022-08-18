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
		tableId := temp[i].CommitTableId
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
func getChain(objectId string) (node TreeNode) {
	temp := NodesTable{}
	Db.Table("nodes").First(&temp, "current_object_id = ?", objectId)
	node.object = ObjectInfo{temp.CurrentObjectId, temp.FatherObjectId, temp.NewConfidence,
		temp.ObjectParameters, temp.ObjectOldLine, temp.ObjectNewLine,
		temp.ObjectDeLine, temp.ObjectAdLine}
	var tempChilds []TreeNode
	Db.Table("nodes").Find(&tempChilds, "father_object_id in (?)", objectId)
	for i := range tempChilds {
		node.childs = append(node.childs, getChain(tempChilds[i].object.objectId))
	}
	return
	//var childs []TreeNode
	//var father ObjectInfo
	//var temp []NodesTable
	//Db.Table("nodes").Find(&temp, "father_object_id in (?)", objectId)
	//node.sonNum = len(temp)
	//if len(temp) == 0{
	//	return
	//}else {
	//	for i := 0; i < len(temp); i++ {
	//		var node1 TreeNode
	//		node1.object.oldObjectId = temp[i].FatherObjectId
	//		node1.object.objectId = temp[i].CurrentObjectId
	//		node1.object.addedLineCount = temp[i].ObjectAdLine
	//		node1.object.deletedlineCount = temp[i].ObjectDeLine
	//		node1.object.newlineCount = temp[i].ObjectNewLine
	//		node1.object.oldlineCount = temp[i].ObjectOldLine
	//		node1.object.parameters =temp[i].ObjectParameters
	//		node1.object.confidence = temp[i].NewConfidence
	//		childs = append(childs, node1)
	//		node1=getChain(childs[i].object.objectId)
	//	}
	//	node.childs = childs
	//	node.object = father
	//}
	//fatherNode := NodesTable{}  //处理第一个传进来的节点
	//Db.Table("nodes").First(&fatherNode,"current_object_id = ?",objectId)
	//father.objectId = fatherNode.CurrentObjectId
	//father.oldObjectId = fatherNode.FatherObjectId
	//father.addedLineCount = fatherNode.ObjectAdLine
	//father.deletedlineCount = fatherNode.ObjectDeLine
	//father.newlineCount = fatherNode.ObjectNewLine
	//father.oldlineCount = fatherNode.ObjectOldLine
	//father.parameters = fatherNode.ObjectParameters
	//father.confidence = fatherNode.NewConfidence
	//node.childs = childs
	//node.object = father
	return
}
