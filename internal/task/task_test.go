package task_test

import (
	"testing"

	"github.com/lemos-g/CtxShift/internal/task"
)

func TestNewCreatesActiveTask(t *testing.T) {
	created, err := task.New("preserve this work")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if got, want := created.Goal(), "preserve this work"; got != want {
		t.Errorf("Goal() = %q, want %q", got, want)
	}

	if got, want := created.State(), task.Active; got != want {
		t.Errorf("State() = %v, want %v", got, want)
	}
}

func TestNewRejectsBlankGoal(t *testing.T) {
	for _, goal := range []string{"", " ", "\t\n"} {
		t.Run("blank", func(t *testing.T) {
			created, err := task.New(goal)
			if err == nil {
				t.Fatal("New() error = nil, want error")
			}
			if created != nil {
				t.Fatalf("New() task = %#v, want nil", created)
			}
		})
	}
}

func TestActiveTaskCanBeCompleted(t *testing.T) {
	created, err := task.New("finish the current task")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := created.Complete(); err != nil {
		t.Fatalf("Complete() error = %v", err)
	}

	if got, want := created.State(), task.Completed; got != want {
		t.Errorf("State() = %v, want %v", got, want)
	}
}

func TestActiveTaskCanBeCancelled(t *testing.T) {
	created, err := task.New("stop this task")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	if err := created.Cancel(); err != nil {
		t.Fatalf("Cancel() error = %v", err)
	}

	if got, want := created.State(), task.Cancelled; got != want {
		t.Errorf("State() = %v, want %v", got, want)
	}
}

func TestCompletedTaskCannotTransition(t *testing.T) {
	for _, transition := range []struct {
		name  string
		apply func(*task.Task) error
	}{
		{"complete", func(task *task.Task) error { return task.Complete() }},
		{"cancel", func(task *task.Task) error { return task.Cancel() }},
	} {
		t.Run(transition.name, func(t *testing.T) {
			created, err := task.New("completed work")
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			if err := created.Complete(); err != nil {
				t.Fatalf("Complete() error = %v", err)
			}

			if err := transition.apply(created); err == nil {
				t.Fatal("transition error = nil, want error")
			}
			if got, want := created.State(), task.Completed; got != want {
				t.Errorf("State() = %v, want %v", got, want)
			}
		})
	}
}

func TestCancelledTaskCannotTransition(t *testing.T) {
	for _, transition := range []struct {
		name  string
		apply func(*task.Task) error
	}{
		{"complete", func(task *task.Task) error { return task.Complete() }},
		{"cancel", func(task *task.Task) error { return task.Cancel() }},
	} {
		t.Run(transition.name, func(t *testing.T) {
			created, err := task.New("cancelled work")
			if err != nil {
				t.Fatalf("New() error = %v", err)
			}
			if err := created.Cancel(); err != nil {
				t.Fatalf("Cancel() error = %v", err)
			}

			if err := transition.apply(created); err == nil {
				t.Fatal("transition error = nil, want error")
			}
			if got, want := created.State(), task.Cancelled; got != want {
				t.Errorf("State() = %v, want %v", got, want)
			}
		})
	}
}
