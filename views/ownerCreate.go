package views

/*import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	. "webService_Refactoring/modules"
)

func OwnerCreate2(context *gin.Context)  {
	var t T3
	err := context.ShouldBind(&t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	pid, err := strconv.Atoi(t.Project.Pid)
	if err != nil {
		context.Status(404)
	}
	release := t.Release
	// report_methods = request.data['methods'] json中的数组传递还为完成解析
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err2 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err2 != nil {
		err2.Error()
	}
	// 数据库中查找pid（未查找到采用json报错）和release中的version未查找到采用json报错）

	context.JSON(http.StatusOK,gin.H{
		"status":"ok",
		"message":"",
		"objects":"",
	})
}
*/
