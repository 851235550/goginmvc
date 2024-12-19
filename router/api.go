package router

import (
	"goginmvc/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	inintJobRouter(r)
}

func inintJobRouter(r *gin.Engine) {
	jobCrtl := controller.NewJobController()

	r.GET("/jobs", jobCrtl.GetJobs)
}
