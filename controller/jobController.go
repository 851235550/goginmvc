package controller

import (
	"goginmvc/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	jobSvc *service.JobService
}

func NewJobController() *JobController {
	return &JobController{
		jobSvc: service.NewJobService(),
	}
}

func (j *JobController) GetJobs(c *gin.Context) {
	c.String(http.StatusOK, "jobs")
}
