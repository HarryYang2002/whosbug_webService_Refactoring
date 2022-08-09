package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api-token-auth", apiToken)

	api := r.Group("/api/v1/users")
	{
		api.POST("/", userCreate)
		api.GET("/:id", userRead)
		api.PUT("/:id", userUpdate)
		api.PATCH("/:id", userPartialUpdate)
	}

	whosbug := r.Group("/whosbug")
	{
		whosbug.POST("/commits/commits-info", commitsInfoCreate)
		whosbug.POST("/commits/delete_uncalculate", commitsDeleteUncalculateCreate)
		whosbug.POST("/commits/diffs", commitsDiffsCreate)
		whosbug.POST("/commits/reviewers", commitsReviewersCreate)
		whosbug.POST("/commits/rules/", commitsRulesCreate)
		whosbug.POST("/commits/train_method", commitsTrainMethodCteate)
		whosbug.POST("/commits.upload-done", commitsUploadDoneCreate)
		whosbug.POST("/create-project-release", createProjectRelease)
		whosbug.POST("/delete_all_related", deleteAllRelated)
		whosbug.GET("/liveness", livenessList)
		whosbug.POST("/owner", OwnerCreate)
		whosbug.POST("/releases/last", releasesLastCreate)

	}
	r.Run(":8083")
}

func apiToken(context *gin.Context) {

}

func userCreate(context *gin.Context) {

}

func userRead(context *gin.Context) {

}

func userUpdate(context *gin.Context) {

}

func userPartialUpdate(context *gin.Context) {

}

func commitsInfoCreate(context *gin.Context) {

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

func createProjectRelease(context *gin.Context) {

}

func deleteAllRelated(context *gin.Context) {

}

func livenessList(context *gin.Context) {

}

func OwnerCreate(context *gin.Context) {

}

func releasesLastCreate(context *gin.Context) {

}
