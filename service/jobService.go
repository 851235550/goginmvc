package service

import "goginmvc/database"

type JobService struct {
	jobManager *database.JobManager
}

func NewJobService() *JobService {
	return &JobService{
		jobManager: database.NewJobManager(),
	}
}

type Job struct {
	ID int `json:"id"`
}

func (s *JobService) GetJobs() ([]*Job, error) {
	jobs := []*Job{}
	dbJob, err := s.jobManager.GetJobs()
	if err != nil {
		return nil, err
	}

	for _, job := range dbJob {
		jobs = append(jobs, &Job{ID: job.ID})
	}

	return jobs, nil
}
