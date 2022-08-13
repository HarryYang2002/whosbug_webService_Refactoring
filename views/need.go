package views

type commitInfo struct {
	commitHash   string
	commitAuthor string
	commitEmail  string
	commitTime   string
}

type objectInfo struct {
	hash             string //Object所属的Commit
	objectId         string //Object的定义链ID
	oldObjectId      string //Object的旧定义链ID
	parameters       string //方法的参数特征
	startLine        int    //起始行
	endLine          int    //结束行
	oldlineCount     int    //旧行数
	newlineCount     int    //新行数
	deletedlineCount int    //移除行数
	addedLineCount   int    //新增行数
}

type historyInfo struct {
	commitHistory commitInfo
	objectHistory objectInfo
}

// [objectPath][author-commitTime]weight
var bugOrigin map[string]map[string]float64

//  @param objectId
//  @return []historyInfo
//  返回的切片要按时间顺序排，最新的commit及其对应object放在索引0
func getHistory(objectId string) (result []historyInfo) {
	return
}

//  @param
//  @return
func getChain() {
	return
}
