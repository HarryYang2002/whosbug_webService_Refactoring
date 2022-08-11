package main

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/views"
)

//包含所有的路由组，go build main.go 即可运行

//TODO 鉴权中间件

func main() {
	r := gin.Default()
	r.POST("/api-token-auth", CreateToken)

	api := r.Group("/api/v1/users")
	{
		api.POST("/", UserCreate)
		r.Use(CheckToken())
		api.GET("/:id", UserRead)
		api.PUT("/:id", UpdateUser)
		api.PATCH("/:id", UpdateUserPartial)
	}

	whosbug := r.Group("/whosbug")
	{
		whosbug.POST("/commits/commits-info", CommitsInfoCreate)
		whosbug.POST("/commits/delete_uncalculate", commitsDeleteUncalculateCreate)
		whosbug.POST("/commits/diffs", commitsDiffsCreate)
		whosbug.POST("/commits/reviewers", commitsReviewersCreate)
		whosbug.POST("/commits/rules/", commitsRulesCreate)
		whosbug.POST("/commits/train_method", commitsTrainMethodCteate)
		whosbug.POST("/commits.upload-done", commitsUploadDoneCreate)
		whosbug.POST("/create-project-release", CreateProjectRelease)
		whosbug.POST("/delete_all_related", deleteAllRelated)
		whosbug.GET("/liveness", LivenessList)
		whosbug.POST("/owner", ownerCreate)
		whosbug.POST("/releases/last", GetLastRelease)

	}
	r.Run(":8083")
}

func commitsDeleteUncalculateCreate(context *gin.Context) {

}

func commitsDiffsCreate(context *gin.Context) {

}

func commitsReviewersCreate(context *gin.Context) {

}

func commitsRulesCreate(context *gin.Context) {

}

func commitsTrainMethodCteate(context *gin.Context) {

}

func commitsUploadDoneCreate(context *gin.Context) {

}

func deleteAllRelated(context *gin.Context) {

}

func ownerCreate(context *gin.Context) {

}

func releasesLastCreate(context *gin.Context) {

}
