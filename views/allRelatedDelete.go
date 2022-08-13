package views

import (
	"github.com/gin-gonic/gin"
)

func AllRelatedDelete(context *gin.Context) {
	//var t T
	//err := context.ShouldBind(&t)
	//if err != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//pid, err2 := strconv.Atoi(t.Project.Pid)
	//if err2 != nil {
	//	context.Status(404)
	//}
	//version := t.Release.Version
	//dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
	//	"sslmode=disable TimeZone=Asia/Shanghai"
	//db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err2 != nil {
	//	err2.Error()
	//}
	//// 通过pid从数据库查找，若错误，返回400，
	//// 再根据pid映射的表单中去查找release.version，若错误，返回400，
	//// err == nil的时候删除有关的所有数据，
	//// 删除失败，返回200和删除失败的信息，成功只返回200
	//temp := DbCreateProject{}
	//res := db.Table("commits").First(&temp, "project_id = ? and release_version = ?", pid, version)
	//if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	//	context.Status(400)
	//	return
	//}
	//temp2 := ObjectType{}
	//tableId := temp.TableId
	//res2 := db.Table("objects").Delete(&temp2, "commit_id = ?", tableId)
	//if res2.Error != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Delete all stuff error1",
	//	})
	//	return
	//}
	//temp3 := UncountedDelete{}
	//res3 := db.Table("uncounted_objects").Delete(&temp3, "release_version = ?", version)
	//if res3.Error != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Delete all stuff error2",
	//	})
	//	return
	//}
	//res4 := db.Table("commits").Delete(&temp, "release_version = ?", version)
	//if res4.Error != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Delete all stuff error3",
	//	})
	//	return
	//}
	////TODO Release表单的删除
	////temp4 := ReleaseType{}
	////res5 := db.Table("releases").Delete(&temp4, "release_version = ?", version)
	////if res5.Error != nil {
	////	context.JSON(http.StatusInternalServerError, gin.H{
	////		"error": "Delete all stuff error3",
	////	})
	////	return
	////}
	//
	//context.JSON(http.StatusOK, gin.H{
	//	"Success": "success",
	//})
}
