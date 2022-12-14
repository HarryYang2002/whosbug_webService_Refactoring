package whosbug

import (
	. "webService_Refactoring/modules"
)

// judge_object
// @param temp2 ObjectsTable
// @param nodes []NodesTable
// @Description uploadDown里面计算时所调用的函数
// @return int
// @return int
// @author: TongLei 2022-08-23 14:47:40
func judge_object(temp2 ObjectsTable, nodes []NodesTable) (int, int) {

	// temp2 := ObjectsTable{}
	// db.Table("objects").Model(&ObjectsTable{}).Find(&temp2)
	//var nodes1 []NodesTable
	//第一次筛选
	var methods []NodesTable
	i := 0
	var nodesnum []int
	var pathsnum []int
	var paramsnum []int
	var tnum int
	for x := 0; x < len(nodes); x++ { //0-9
		if nodes[x].CurrentObjectID == temp2.CurrentObjectID {
			methods = append(methods, nodes[x])
			nodesnum = append(nodesnum, x) //3 5 7 9 n=4
		} //0 1 2 3
	}
	if len(methods) == 0 {
		return 0, 0
	}
	//第二次筛选
	var path []NodesTable
	for x := 0; x < len(methods); x++ {
		if methods[x].ObjectPath == temp2.ObjectPath {
			path = append(path, methods[x])
			pathsnum = append(pathsnum, x) //1 3 n=2  3  9
		} //0 1
	}
	if len(path) == 0 {
		return 0, 0
	}
	//第三次筛选
	var params []NodesTable
	for x := 0; x < len(path); x++ {
		if path[x].ObjectParameters == temp2.Parameters {
			params = append(params, path[x])
			paramsnum = append(paramsnum, x) //1  9
			i = x
		}
	}
	if len(params) == 0 {
		return 0, 0
	} else {
		tnum = nodesnum[pathsnum[i]]
	}

	return i, tnum

}

func judge_change(object UncalculateObjectInfo) int {
	if object.addedLineCount == 0 && object.deletedlineCount == 0 {
		return 1
	} else {
		return 0
	}

}
