package database

import (
	"github.com/jmoiron/sqlx"
)

type JobManager struct {
	db *sqlx.DB
}

func NewJobManager() *JobManager {
	return &JobManager{
		db: dbInstance,
	}
}

type Job struct {
	ID int `db:"id"`
}

func (j *JobManager) GetJobs() ([]*Job, error) {
	jobs := []*Job{}

	err := j.db.Select(&jobs, "SELECT * FROM jobs")

	if err != nil {
		return nil, err
	}
	return jobs, nil
}
