package middlewear

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
	. "webService_Refactoring/modules"
)

// CheckToken 检查token
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "Authentication credentials were not provided.",
			})
			c.Abort()
			return
		}
		arr := strings.Fields(token)
		realToken := arr[1]
		temp := UsersTable{}
		res := Db.Table("users").Where("user_token = ?", realToken).First(&temp)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"detail": "Authentication credentials were not provided.",
			})
			c.Abort()
			return
		}
	}
}
