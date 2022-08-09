package views

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

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
