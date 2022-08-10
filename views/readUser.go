package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "webService_Refactoring/modules"
)

func UserRead(context *gin.Context) {
	var user UserID
	err := context.ShouldBindUri(&user)
	if err != nil {
		context.Status(400)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"username":   "",
		"first_name": "",
		"last_name":  "",
	})
}
