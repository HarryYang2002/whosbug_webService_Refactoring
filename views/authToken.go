package views

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	. "webService_Refactoring/modules"
)

// MD5 md5算法，生成的依据是时间戳
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// CreateToken 路由函数，访问端口时调用，根据用户输入的username和password，经过MD5算法后生成token
// 并将username、password、token、和根据时间戳生成的uuid存入数据库中
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

	dbCreateUser := DbCreateUser{
		UserId:       CreateUUID(),
		UserName:     loginForm.Username,
		UserToken:    MD5(token),
		UserPassword: loginForm.Password,
	}

	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	fmt.Println(db.Table("users").Create(&dbCreateUser).RowsAffected)

}

// CreateUUID 根据时间戳生成uuid
func CreateUUID() uuid.UUID {
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	return u1
}
