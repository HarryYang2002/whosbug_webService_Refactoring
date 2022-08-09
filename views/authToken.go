package views

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

// Claims 是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	username string `json:"username" binding:"required"`
	password string `json:"password" binding:"required"`
	jwt.StandardClaims
}

func CreateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var claims Claims
		err := context.ShouldBind(&claims)
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				context.JSON(http.StatusOK, gin.H{
					"msg": errs.Error(),
				})
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"error": errs.Error(),
			})
			return
		}
		token, err := GenerateToken(claims.username, claims.password)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

var jwtSecret = []byte("setting.JwtSecret")

// GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username: username,
		password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
