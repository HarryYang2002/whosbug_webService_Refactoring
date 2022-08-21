package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	token2 "webService_Refactoring/api/v1/token"
	. "webService_Refactoring/modules"
)

// UserCreate 生成用户数据并存储到数据库中
func UserCreate(context *gin.Context) {
	//注册命名规则
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
	token := token2.MD5(tokenKey)

	var userUUID uuid.UUID
	userUUID = token2.CreateUUID()

	context.JSON(http.StatusOK, gin.H{
		"id":         userUUID.String(),
		"username":   registerForm.Username,
		"password":   registerForm.Password,
		"first_name": registerForm.Firstname,
		"last_name":  registerForm.Lastname,
		"email":      registerForm.Email,
		"auth_token": token,
	})

	DbCreateUser := UsersTable{
		UserID:        userUUID,
		UserName:      registerForm.Username,
		UserToken:     token,
		UserPassword:  token2.MD5(registerForm.Password),
		UserFirstName: registerForm.Firstname,
		UserLastName:  registerForm.Lastname,
		UserEmail:     registerForm.Email,
	}
	Db.Table("users").Create(&DbCreateUser)
}
