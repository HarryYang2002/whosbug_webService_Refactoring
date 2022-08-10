package views

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"webService_Refactoring/modules"
)

func UserCreate(context *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("usernamerule", modules.UsernameRule)
	}
	var registerForm modules.RegisterForm
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

	context.JSON(http.StatusOK, gin.H{
		"id":           "", //数据库中的id
		"username":     registerForm.Username,
		"password":     registerForm.Password,
		"first_name":   registerForm.Firstname,
		"last_name":    registerForm.Lastname,
		"email":        registerForm.Email,
		"authon_token": token,
	})
}
