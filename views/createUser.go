package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	. "webService_Refactoring/modules"
)

func UserCreate(context *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("usernamerule", UsernameRule)
	}
	var registerForm RegisterForm
	err := context.ShouldBind(&registerForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var tokenKey string
	tokenKey = registerForm.Username + "&" + registerForm.Password
	token := MD5(tokenKey)

	var userUUID uuid.UUID
	userUUID = CreateUUID()

	context.JSON(http.StatusOK, gin.H{
		"id":         userUUID.String(),
		"username":   registerForm.Username,
		"password":   registerForm.Password,
		"first_name": registerForm.Firstname,
		"last_name":  registerForm.Lastname,
		"email":      registerForm.Email,
		"auth_token": token,
	})

	dbCreateUser := DbCreateUser{
		UserId:        userUUID,
		UserName:      registerForm.Username,
		UserToken:     token,
		UserPassword:  registerForm.Password,
		UserFirstName: registerForm.Firstname,
		UserLastName:  registerForm.Lastname,
		UserEmail:     registerForm.Email,
	}
	dsn := "host=localhost user=postgres password=123456 dbname=whobug2022 port=5433 " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	fmt.Println(db.Table("users").Create(&dbCreateUser).RowsAffected)
}
