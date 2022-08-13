package views

import (
	"github.com/gin-gonic/gin"
)

func CommitsTrainMethodCreate(context *gin.Context) {
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
	//lastCommitHash := t.Release.CommitHash
	//dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
	//	"sslmode=disable TimeZone=Asia/Shanghai"
	//db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err2 != nil {
	//	err2.Error()
	//}
	//// 通过pid从数据库查找，若错误，返回400，
	//// 再去查找release.version，若错误，返回400，
	//// err == nil的时候删除uncalculate的数据，
	//// 删除失败，返回200和删除失败的信息，成功只返回200
	//temp := DbCreateProject{}
	//res := db.Table("commits").First(&temp, "project_id = ? ", pid)
	//if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	//	context.Status(400)
	//	return
	//}
	//errs := db.Table("commits").First(&temp, "release_version = ?and pv_last_commit_hash =", version, lastCommitHash)
	//if errs != nil {
	//	context.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Delete error",
	//	})
	//	return
	//}
	//context.Status(200)
	//count := 100
	//
	//// 创建进度条并开始
	//bar := pb.StartNew(count)
	//
	//for i := 0; i < count; i++ {
	//	bar.Increment()
	//	time.Sleep(50 * time.Microsecond)
	//}
	//
	//// 结束进度条
	//bar.Finish()
}
