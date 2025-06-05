package task

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

type TaskStatus string

const (
	Pending   TaskStatus = "pending"
	Active    TaskStatus = "active"
	Completed TaskStatus = "completed"
	Failed    TaskStatus = "failed"
)

var (
	ErrTaskAlreadyClaimed = errors.New("task-already-claimed")
)

type Manager interface {
	Create(ctx context.Context, id string, task Task) error
	Claim(ctx context.Context, id string, assignee string) error
	Complete(ctx context.Context, id string, output string, err error) error
	Fail(ctx context.Context, id string, err error) error
}

type Task struct {
	ID        string     `json:"id"`
	Assignee  string     `json:"assignee"`
	Client    string     `json:"client"`
	Input     string     `json:"input"`
	Output    string     `json:"output"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewTask(id string, client string, input string) Task {
	return Task{
		ID:        id,
		Client:    client,
		Input:     input,
		Status:    Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t Task) JSONMarshal() []byte {
	b, _ := json.Marshal(t)
	return b
}

func (t *Task) JSONUnmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}
