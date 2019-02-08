package main

import (
	"time"

	"github.com/google/uuid"
)

// Job is
type Job struct {
	ID             string `gorm:"type:varchar(40); primary_key"`
	MaintainerUUID string `gorm:"type:varchar(40)"`
	Public         bool   `gorm:"default: false; not null"`
	Backend        string `gorm:"varchar(255)"`
	SourceURL      string `gorm:"type:text; not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewJobOption is a struct of essential fields to create new job.
// All fields required.
type NewJobOption struct {
	MaintainerUUID string
	Public         bool
	Backend        string
	SourceURL      string
}

// NewJob create new job, save to database, and return it.
func NewJob(opt *NewJobOption) (*Job, error) {
	id, _ := uuid.NewUUID()
	job := Job{
		ID:             id.String(),
		MaintainerUUID: opt.MaintainerUUID,
		Public:         opt.Public,
		Backend:        opt.Backend,
		SourceURL:      opt.SourceURL,
	}

	// TODO - save job to database

	return &job, nil
}

// Execute function create new execution of the job, and start it.
func (job *Job) Execute(triggeredBy string) (*Execution, error) {
	e := NewExecution(&NewExecutionOption{
		JobID:       job.ID,
		TriggeredBy: triggeredBy,
	})
	return e, e.Start()
}
