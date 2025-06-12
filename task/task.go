package task

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/gofrs/uuid/v5"
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
	RewardDefault         = big.NewInt(150000).String()
)

type Manager interface {
	Create(ctx context.Context, id string, task Task) error
	Claim(ctx context.Context, id string, assignee string) error
	Complete(ctx context.Context, id string, output string, err error) error
	Fail(ctx context.Context, id string, err error) error
}

type Task struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Assignee    string     `json:"assignee"`
	Input       string     `json:"input"`
	Output      string     `json:"output"`
	Reward      string     `json:"reward"`
	Status      TaskStatus `json:"status"`
	Submitter   string     `json:"submitter"`
	Transaction string     `json:"transaction"`
	Contract    string     `json:"contract"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(client, submitter, input string) Task {
	return Task{
		ID:        uuid.Must(uuid.NewV4()).String(),
		UserID:    client,
		Submitter: submitter,
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
