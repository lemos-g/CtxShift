package task

import (
	"errors"
	"strings"
)

var (
	errEmptyGoal     = errors.New("task goal must not be empty")
	errTaskNotActive = errors.New("task is not active")
)

type State uint8

const (
	Active State = iota
	Completed
	Cancelled
)

type Task struct {
	goal  string
	state State
}

func New(goal string) (*Task, error) {
	if strings.TrimSpace(goal) == "" {
		return nil, errEmptyGoal
	}

	return &Task{
		goal:  goal,
		state: Active,
	}, nil
}

func (t *Task) Goal() string {
	return t.goal
}

func (t *Task) State() State {
	return t.state
}

func (t *Task) Complete() error {
	if t.state != Active {
		return errTaskNotActive
	}

	t.state = Completed
	return nil
}

func (t *Task) Cancel() error {
	if t.state != Active {
		return errTaskNotActive
	}

	t.state = Cancelled
	return nil
}
