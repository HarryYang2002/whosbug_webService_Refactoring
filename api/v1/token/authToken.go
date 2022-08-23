package token

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	. "webService_Refactoring/modules"
)

// MD5
// @param v string
// @Description md5算法，生成的依据是时间戳
// @return string
// @author: HarryYang 2022-08-23 15:44:44
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// CreateToken
// @param context *gin.Context
// @Description 据用户输入的username和password，经过MD5算法后生成token
// @author: HarryYang 2022-08-23 15:44:15
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

// CreateUUID
// @Description 根据时间戳生成uuid
// @return uuid.UUID
// @author: HarryYang 2022-08-23 15:44:56
func CreateUUID() uuid.UUID {
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return u1
}
