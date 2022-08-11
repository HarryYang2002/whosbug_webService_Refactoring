package views

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	. "webService_Refactoring/modules"
)

func CommitsInfoCreate(context *gin.Context) {

	var t T2

	//err3 := context.ShouldBind(&t)

	// if err3 != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err3.Error(),
	// 	})
	// 	return
	// }
	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	err := json.Unmarshal(bodyBytes, &t)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{

		// "project": t.Project,
		// "release": t.Release,
		"": t,
	})
}
