package views

import (
	"github.com/gin-gonic/gin"
)

func GetLastRelease(c *gin.Context) {
	//var id modules.ProjectID
	//if err := c.ShouldBind(&id); err != nil {
	//	err.Error()
	//}
	//dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
	//	"sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	err.Error()
	//}
	//temp := modules.DbCreateProject{}
	//res := db.Table("commits").Where("project_id = ?", id.Pid).First(&temp)
	//if errors.Is(res.Error, gorm.ErrRecordNotFound) {
	//	info := "The project with pid = " + strconv.Itoa(id.Pid) + " does not exists."
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"error": info,
	//	})
	//	return
	//}
	//c.JSON(http.StatusCreated, gin.H{
	//	"lastReleaseVersion": temp.ReleaseVersion,
	//})

}
