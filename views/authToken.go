package views

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	. "webService_Refactoring/modules"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func CreateToken(context *gin.Context) {
	var loginForm LoginForm
	err := context.ShouldBind(&loginForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var token string
	token = loginForm.Username + "&" + loginForm.Password
	context.JSON(http.StatusOK, gin.H{
		"token": MD5(token),
	})
}
