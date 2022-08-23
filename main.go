package main

import (
	_ "net/http/pprof"
	. "webService_Refactoring/modules"
	. "webService_Refactoring/routes"
)

// main
// @Description 入口函数
// @author: WuTianPeng 2022-08-23 15:50:25
func main() {

	// 引用数据库
	InitDB()

	// 启动路由
	InitRouter()
}
