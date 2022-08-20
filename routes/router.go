package routes

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/api/v1/token"
	. "webService_Refactoring/api/v1/users"
	. "webService_Refactoring/api/v1/whosbug"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
)

func InitRouter() {
	gin.SetMode(AppMode)
	r := gin.Default()
	r.POST("/v1/api-token-auth", CreateToken)

	api := r.Group("/v1/users")
	{
		api.POST("/", UserCreate)
		api.GET("/:id", CheckToken(), UserRead)
		api.PUT("/:id", CheckToken(), UpdateUser)
		api.PATCH("/:id", CheckToken(), UpdateUserPartial)
	}

	commits := r.Group("/v1/commits")
	{
		commits.POST("/commits-info", CheckToken(), CommitsInfoCreate) //1
		//commits.POST("/delete_uncalculate", CheckToken(), UncalculateDelete) //1
		commits.POST("/diffs", CheckToken(), CommitsDiffsCreate) //1
		//review 暂时不重构
		commits.POST("/reviewers", CheckToken(), CommitsReviewersCreate)
		commits.POST("/rules/", CheckToken(), CommitsRulesCreate)
		//
		commits.POST("/train_method", CommitsTrainMethodCreate, CheckToken()) //1
		commits.POST("/upload-done", CheckToken(), CommitsUploadDoneCreate)   //1
	}
	r.POST("/v1/create-project-release", CheckToken(), CreateProjectRelease) //1
	r.POST("/v1/delete_all_related", CheckToken(), AllRelatedDelete)         //1
	r.GET("/v1/liveness", CheckToken(), LivenessList)                        //1
	r.POST("/v1/owner", CheckToken(), OwnerCreate)                           //1
	r.POST("/v1/releases/last", CheckToken(), GetLastRelease)                //1
	r.Run(HttpPort)
}
