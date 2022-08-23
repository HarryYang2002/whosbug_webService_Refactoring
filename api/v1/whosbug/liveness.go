package whosbug

import "github.com/gin-gonic/gin"

// LivenessList
// @param context *gin.Context
// @Description secure the liveness of the webservice
// @author: HarryYang 2022-08-23 14:45:04
func LivenessList(context *gin.Context) {
	context.Status(200)
}
