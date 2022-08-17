package views

import (
	"fmt"
	. "webService_Refactoring/modules"
)

func judge_object(temp2 ObjectsTable, nodes []NodesTable) (int, []NodesTable, int) {

	// temp2 := ObjectsTable{}
	// db.Table("objects").Model(&ObjectsTable{}).Find(&temp2)
	//var nodes1 []NodesTable
	//第一次筛选
	var methods []NodesTable
	var i int
	var nodesnum []int
	var pathsnum []int
	var paramsnum []int
	var tnum int
	for x := 0; x < len(nodes); x++ {
		if nodes[x].CurrentObjectId == temp2.CurrentObjectId {
			methods = append(methods, nodes[x])
			nodesnum = append(nodesnum, x) //3 5 7 9
		}
	}
	if len(methods) == 0 {
		fmt.Println("Get objects error:")
		return 0, methods, 0
	}
	//第二次筛选
	var path []NodesTable
	for x := 0; x < len(methods); x++ {
		if methods[x].ObjectPath == temp2.ObjectPath {
			path = append(path, nodes[x])
			pathsnum = append(pathsnum, x) //1 3
		}
	}
	if len(path) == 0 {
		fmt.Println("Get objects error:")
		return 0, path, 0

	}
	//第三次筛选
	var params []NodesTable
	for x := 0; x < len(path); x++ {
		if path[x].ObjectParameters == temp2.Parameters {
			params = append(params, path[x])
			paramsnum = append(paramsnum, x) //1
			i = x
		}
	}
	if len(params) == 0 {
		fmt.Println("Get objects error:")
		return 0, params, 0
	} else {
		tnum = nodesnum[pathsnum[i]]
	}

	return i, params, tnum

}

func judge_change(object UncalculateObjectInfo) int {
	if object.addedLineCount == 0 && object.deletedlineCount == 0 {
		return 1
	} else {
		return 0
	}
}
