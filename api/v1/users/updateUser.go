package users

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

// UpdateUser
// @param context *gin.Context
// @Description 更新用户信息，put为上传
// @author: HarryYang 2022-08-23 15:41:58
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

	temp := UsersTable{}
	var searchId string
	searchId = context.Param("id")
	fmt.Println(searchId)
	//put方法决定以form-data进行传递数据

	fn := context.PostForm("first_name")
	ln := context.PostForm("last_name")

	res := Db.Table("users").First(&temp, "user_id = ?", searchId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = fn
	temp.UserLastName = ln

	er := Db.Table("users").Where("user_id = ?", searchId).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}

// UpdateUserPartial
// @param context *gin.Context
// @Description 更新用户信息，patch为修改
// @author: HarryYang 2022-08-23 15:42:11
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

	temp := UsersTable{}
	var searchId string
	searchId = context.Param("id")
	fmt.Println(searchId)

	newfn := context.PostForm("first_name")
	newln := context.PostForm("last_name")

	res := Db.Table("users").First(&temp, "user_id = ?", searchId)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		context.Status(401)
		return
	}
	temp.UserFirstName = newfn
	temp.UserLastName = newln

	er := Db.Table("users").Where("user_id = ?", searchId).Updates(&temp).Error
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         temp.UserID,
		"username":   temp.UserName,
		"first_name": temp.UserFirstName,
		"last_name":  temp.UserLastName,
	})
}
