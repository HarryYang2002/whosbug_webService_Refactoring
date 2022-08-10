package views

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	temp := DbCreateUser{}
	var searchId string
	searchId = context.Param("id")
	fmt.Println(searchId)
	//searchId, _ := uuid.Parse(Id)

	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	//first为查询，可以返回查询错误，Find不能返回错误
	res := db.Table("users").First(&temp, "user_id = ?", searchId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserId,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})

}
