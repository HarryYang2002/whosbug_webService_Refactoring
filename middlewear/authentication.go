package middlewear

import "github.com/gin-gonic/gin"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Get("Authorization")
	}
}
