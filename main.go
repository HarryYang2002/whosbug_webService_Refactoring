package main

import (
	"github.com/gin-gonic/gin"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/views"
)

//包含所有的路由组，go build main.go 即可运行

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
		whosbug.POST("/commits/delete_uncalculate", UncalculateDelete)
		whosbug.POST("/commits/diffs", CommitsDiffsCreate)
		//review 暂时不重构
		whosbug.POST("/commits/reviewers", CommitsReviewersCreate)
		whosbug.POST("/commits/rules/", CommitsRulesCreate)
		//
		whosbug.POST("/commits/train_method", CommitsTrainMethodCreate)
		whosbug.POST("/commits.upload-done", CommitsUploadDoneCreate)
		whosbug.POST("/create-project-release", CreateProjectRelease)
		whosbug.POST("/delete_all_related", AllRelatedDelete)
		whosbug.GET("/liveness", LivenessList)
		whosbug.POST("/owner", OwnerCreate) // 新功能接口（接受堆栈信息，数据库匹配，返回需计算函数的切片）
		whosbug.POST("/releases/last", GetLastRelease)
	}
	r.Run(":8083")
}
