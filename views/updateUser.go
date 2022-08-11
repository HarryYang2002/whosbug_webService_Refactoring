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

	temp := DbCreateUser{}
	var searchId string
	searchId = context.Param("id")
	fmt.Println(searchId)
	//put方法决定以form-data进行传递数据

	fn := context.PostForm("first_name")
	ln := context.PostForm("last_name")

	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	res := db.Table("users").First(&temp, "user_id = ?", searchId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = fn
	temp.UserLastName = ln

	er := db.Table("users").Where("user_id = ?", searchId).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserId,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}

func UpdateUserPartial(context *gin.Context) {
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

	temp := DbCreateUser{}
	var searchId string
	searchId = context.Param("id")
	fmt.Println(searchId)

	newfn := context.PostForm("first_name")
	newln := context.PostForm("last_name")

	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	res := db.Table("users").First(&temp, "user_id = ?", searchId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = newfn
	temp.UserLastName = newln

	er := db.Table("users").Where("user_id = ?", searchId).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserId,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}
