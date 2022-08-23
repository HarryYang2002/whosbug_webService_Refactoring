package whosbug

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// AllRelatedDelete
// @param context *gin.Context
// @Description 删除commits、objects、releases表中的数据
// @author: HarryYang 2022-08-23 14:25:17
func AllRelatedDelete(context *gin.Context) {
	var t T
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//提取pid、version
	pid := t.Project.Pid
	version := t.Release.Version
	//以pid去找
	project := ProjectsTable{}
	res := Db.Table("projects").Where("project_id = ?", pid).First(&project)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Project get fails",
			"detail": "no such project:" + pid,
		})
		return
	}
	//以version去找
	release := ReleasesTable{}
	res2 := Db.Table("releases").Where("release_version = ? and project_table_id = ?", version, project.TableID).First(&release)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":  "Release get fails",
			"detail": "no such release:" + version,
		})
		return
	}
	//删除的内容在此做一下说明
	//首先去releases表中去找对应的version，取出该条数据的table_id,此时可删除该条数据
	//再去commits表中去找table_id对应的release_table_id的那条数据,此时可删除该条数据
	//再以该条数据的table_id去objects表中相应的commit_table_id，删除
	//后续数据库设置了级联删除，可以优化
	realRelease := ReleasesTable{}
	uncounted := ObjectsTable{}
	commit := CommitsTable{}
	Db.Table("releases").First(&realRelease, "release_version = ?", version)
	releaseId := realRelease.TableID
	res3 := Db.Table("releases").Delete(&realRelease, "release_version = ?", version)
	if res3.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete all stuff error",
		})
		return
	}
	Db.Table("commits").First(&commit, "release_table_id = ?", releaseId)
	uncountedId := commit.TableID
	res4 := Db.Table("commits").Delete(&realRelease, "release_table_id = ?", releaseId)
	if res4.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete all stuff error",
		})
		return
	}
	res5 := Db.Table("objects").Delete(&uncounted, "commit_table_id = ?", uncountedId)
	if res5.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete all stuff error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"Success": "Success",
	})
}
