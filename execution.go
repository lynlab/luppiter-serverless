package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/qor/transition"
)

// Execution is a object for each trigger of jobs.
type Execution struct {
	ID          string `gorm:"type:varchar(40); primary_key"`
	JobID       string `gorm:"type:varchar(40); not null"`
	Job         Job
	TriggeredBy string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	transition.Transition
}

// NewExecutionOption is a struct of essential fields to create new execution.
// All fields required.
type NewExecutionOption struct {
	JobID       string
	TriggeredBy string
}

// States:
//   - created
//   - started  <- created
//   - canceled <- created, started
//   - finished <- started
var executionSM *transition.StateMachine

// NewExecution create new execution object and return it.
func NewExecution(opt *NewExecutionOption) *Execution {
	id, _ := uuid.NewUUID()
	exe := Execution{
		ID:    id.String(),
		JobID: opt.JobID,
	}
	exe.SetState("created")

	// TODO - save job to database

	return &exe
}

// Start starts execution.
func (e *Execution) Start() error {
	return executionSM.Trigger("start", e, nil)
}

func init() {
	executionSM = transition.New(&Execution{})
	executionSM.Initial("created")
	executionSM.State("started")
	executionSM.State("canceled")
	executionSM.State("finished")

	executionSM.Event("start").To("started").From("created")
}
