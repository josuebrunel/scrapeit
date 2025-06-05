package task

import (
	"context"
	"scrapeit/pkg/env"
	"strings"
	"time"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	rdb *redis.Client
}

func NewRedisManager() *RedisManager {
	url := env.Get(env.RedisAddrN)
	xlog.Info("Redis URL", "url", url)
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		xlog.Error("Error while connecting to redis", "error", err)
		panic(err)
	}
	return &RedisManager{
		rdb: rdb,
	}
}

func (m *RedisManager) Close() error {
	return m.rdb.Close()
}

func (m RedisManager) Create(ctx context.Context, key string, task Task) error {
	err := m.rdb.Set(ctx, key, task.JSONMarshal(), 0).Err()
	if err != nil {
		xlog.Error("Error while creating task", "id", key, "error", err)
		return err
	}
	return nil
}

func (m RedisManager) Get(ctx context.Context, key string) (Task, error) {
	var task Task
	data, err := m.rdb.Get(ctx, key).Bytes()
	if err != nil {
		xlog.Error("Error while getting task", "id", key, "error", err)
		return task, err
	}
	err = task.JSONUnmarshal(data)
	if err != nil {
		xlog.Error("Error while unmarshalling task", "id", key, "error", err)
		return task, err
	}
	return task, nil
}

func (m RedisManager) Update(ctx context.Context, key string, task Task) error {
	task.UpdatedAt = time.Now()
	err := m.rdb.Set(ctx, key, task.JSONMarshal(), 0).Err()
	if err != nil {
		xlog.Error("Error while updating task", "id", key, "error", err)
		return err
	}
	return nil
}

func (m RedisManager) Delete(ctx context.Context, key string) error {
	err := m.rdb.Del(ctx, key).Err()
	if err != nil {
		xlog.Error("Error while deleting task", "id", key, "error", err)
		return err
	}
	return nil
}

func (m RedisManager) Claim(ctx context.Context, key string, assignee string) error {
	task, err := m.Get(ctx, key)
	if err != nil {
		return err
	}
	if !strings.EqualFold(task.Assignee, "") || task.Status != Pending {
		xlog.Warn("Task already claimed", "id", key, "task", task)
		return ErrTaskAlreadyClaimed
	}
	task.Assignee = assignee
	task.Status = Active
	task.UpdatedAt = time.Now()
	err = m.Update(ctx, key, task)
	if err != nil {
		return err
	}
	return nil
}

func (m RedisManager) Complete(ctx context.Context, key string, output string, jobErr error) error {
	task, err := m.Get(ctx, key)
	if err != nil {
		return err
	}
	if jobErr != nil {
		task.Status = Failed
		task.Output = jobErr.Error()
		task.UpdatedAt = time.Now()
	} else {
		task.Status = Completed
	}
	task.Output = output
	task.UpdatedAt = time.Now()
	err = m.Update(ctx, key, task)
	if err != nil {
		return err
	}
	return nil
}
