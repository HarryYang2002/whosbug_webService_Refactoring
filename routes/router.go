package routes

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/api/v1/token"
	. "webService_Refactoring/api/v1/users"
	. "webService_Refactoring/api/v1/whosbug"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
)

// InitRouter
// @Description 总路由
// @author: HarryYang 2022-08-23 14:29:33
func InitRouter() {
	gin.SetMode(AppMode)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/token", CreateToken)
		api := v1.Group("/users")
		{
			api.POST("/", UserCreate)
			api.GET("/:id", CheckToken(), UserRead)
			api.PUT("/:id", CheckToken(), UpdateUser)
			api.PATCH("/:id", CheckToken(), UpdateUserPartial)
		}
		commits := v1.Group("/commits")
		{
			commits.POST("/commits_info", CheckToken(), CommitsInfoCreate)
			commits.POST("/uncalculate_delete", CheckToken(), UncalculateDelete)
			commits.POST("/diffs", CheckToken(), CommitsDiffsCreate)
			commits.POST("/train_method", CheckToken(), CommitsTrainMethodCreate)
			commits.POST("/nodes_create", CheckToken(), CommitsUploadDoneCreate)
			//reviews、rules 暂时不重构
			commits.POST("/reviewers", CheckToken(), CommitsReviewersCreate)
			commits.POST("/rules", CheckToken(), CommitsRulesCreate)
		}
		v1.POST("/project_release_create", CheckToken(), CreateProjectRelease)
		v1.POST("/all_related_delete", CheckToken(), AllRelatedDelete)
		v1.GET("/liveness", CheckToken(), LivenessList)
		v1.POST("/owner", CheckToken(), OwnerCreate)
		v1.POST("/last_releases", CheckToken(), GetLastRelease)
	}
	r.Run(HttpPort)
}
