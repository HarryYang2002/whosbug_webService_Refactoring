package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "webService_Refactoring/modules"
)

func UpdateUser(context *gin.Context) {
	var userId UserID
	err := context.ShouldBindUri(&userId)
	if err != nil {
		context.Status(400)
		return
	}
	var updateUser UpdateUsers
	errs := context.ShouldBind(&updateUser)
	if errs != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":       userId.ID,
		"username": "",
		// 暂定直接传递，后期连接数据库后由数据库提供，保证更改已经上传到数据库
		"first_name": updateUser.FirstName,
		"last_name":  updateUser.LastName,
	})
}

//func UpdateUserPartial(context *gin.Context) {
//	var userId modules.UserID
//	err := context.ShouldBindUri(&userId)
//	if err != nil {
//		context.Status(400)
//		return
//	}
//	var updateUser modules.UpdateUsers
//	errs := context.ShouldBind(&updateUser)
//	if errs != nil {
//		context.Status(400)
//		return
//	}
//	context.JSON(http.StatusOK, gin.H{
//		"id":         userId.ID,
//		"username":   "",
//		"first_name": updateUser.FirstName,
//		"last_name":  updateUser.LastName,
//	})
//}
