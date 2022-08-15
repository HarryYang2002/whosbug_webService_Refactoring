package views

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
	return
}

//  @param objectId 函数的id
//  @return	chainNode 该函数所在的定义链的根结点
func getChain(objectId string) (node TreeNode) {
	return node
}
