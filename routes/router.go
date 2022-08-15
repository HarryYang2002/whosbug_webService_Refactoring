package routes

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
	. "webService_Refactoring/views"
)

func InitRouter() {
	gin.SetMode(AppMode)
	r := gin.Default()
	r.POST("/api-token-auth", CreateToken)

	api := r.Group("/v1/users")
	{
		api.POST("/", UserCreate)
		r.Use(CheckToken())
		api.GET("/:id", UserRead)
		api.PUT("/:id", UpdateUser)
		api.PATCH("/:id", UpdateUserPartial)
	}

	commits := r.Group("/commits")
	{
		commits.POST("/commits-info", CommitsInfoCreate)
		commits.POST("/delete_uncalculate", UncalculateDelete)
		commits.POST("/diffs", CommitsDiffsCreate)
		//review 暂时不重构
		commits.POST("/reviewers", CommitsReviewersCreate)
		commits.POST("/rules/", CommitsRulesCreate)
		//
		commits.POST("/train_method", CommitsTrainMethodCreate)
		commits.POST(".upload-done", CommitsUploadDoneCreate)
	}
	r.POST("create-project-release", CreateProjectRelease)
	r.POST("delete_all_related", AllRelatedDelete)
	r.GET("liveness", LivenessList)
	r.POST("owner", OwnerCreate)
	r.POST("releases/last", GetLastRelease)
	r.Run(HttpPort)
}
