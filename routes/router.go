package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"
	. "webService_Refactoring/middlewear"
	. "webService_Refactoring/utils"
	. "webService_Refactoring/views"
)

var srv *http.Server

func InitRouter() {
	file, _ := os.Create("./cpufile.prof")
	pprof.StartCPUProfile(file)
	gin.SetMode(AppMode)
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":8083",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	r.POST("/v1/api-token-auth", CreateToken)

	api := r.Group("/v1/users")

	api.POST("/", UserCreate)
	r.Use(CheckToken())
	api.GET("/:id", UserRead)
	api.PUT("/:id", UpdateUser)
	api.PATCH("/:id", UpdateUserPartial)

	commits := r.Group("/v1/commits")

	commits.POST("/commits-info", CommitsInfoCreate)       //1
	commits.POST("/delete_uncalculate", UncalculateDelete) //1
	commits.POST("/diffs", CommitsDiffsCreate)             //1
	//review 暂时不重构
	commits.POST("/reviewers", CommitsReviewersCreate)
	commits.POST("/rules/", CommitsRulesCreate)
	//
	commits.POST("/train_method", CommitsTrainMethodCreate) //1
	commits.POST("/upload-done", CommitsUploadDoneCreate)   //1

	r.POST("/v1/create-project-release", CreateProjectRelease) //1
	r.POST("/v1/delete_all_related", AllRelatedDelete)         //1
	r.GET("/v1/liveness", LivenessList)                        //1
	r.POST("/v1/owner", OwnerCreate)                           //1
	r.POST("/v1/releases/last", GetLastRelease)                //1

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer file.Close()
	defer pprof.StopCPUProfile()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	// r.Run(HttpPort)
}
