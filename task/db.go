package task

import (
	"context"
	"scrapeit/pkg/models"

	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

func (m EtherManager) DBGet(ctx context.Context, id string) (Task, error) {
	tsk, err := models.FindTask(ctx, m.db, uuid.Must(uuid.FromString(id)))
	if err != nil {
		xlog.Error("Failed to find task", "error", err)
		return Task{}, err
	}
	return Task{
		ID:          tsk.ID.String(),
		Assignee:    tsk.Assignee.V,
		Submitter:   tsk.Submitter.V,
		Contract:    tsk.Contract.V,
		Transaction: tsk.Transaction.V,
		Input:       tsk.Input.V,
		Output:      tsk.Output.V,
		Status:      TaskStatus(tsk.Status.V),
		CreatedAt:   tsk.CreatedAt.V,
		UpdatedAt:   tsk.UpdatedAt.V,
	}, nil
}

func (m EtherManager) Upsert(ctx context.Context, tsk Task) error {
	var (
		id  = uuid.Must(uuid.FromString(tsk.ID))
		err error
		uid uuid.UUID
	)
	if tsk.UserID != "" {
		uid = uuid.Must(uuid.FromString(tsk.UserID))
	}
	xlog.Debug("Upserting task", "task", tsk)
	_, err = models.Tasks.Insert(&models.TaskSetter{
		ID:          &id,
		UserID:      &uid,
		Assignee:    models.NullFrom(tsk.Assignee),
		Contract:    models.NullFrom(tsk.Contract),
		Input:       models.NullFrom(tsk.Input),
		Output:      models.NullFrom(tsk.Output),
		Status:      models.NullFrom(string(tsk.Status)),
		Submitter:   models.NullFrom(tsk.Submitter),
		Transaction: models.NullFrom(tsk.Transaction),
		UpdatedAt:   models.NullFrom(tsk.UpdatedAt),
	},
		im.OnConflict("id").DoUpdate(
			im.SetExcluded("assignee", "output", "status", "transaction", "updated_at"),
		),
	).One(ctx, m.db)
	if err != nil {
		xlog.Error("Failed to convert user ID to UUID", "error", err)
		return err
	}
	return nil
}

func QueryListByUser(ctx context.Context, db bob.DB, userID string, limit int) (models.TaskSlice, error) {
	uid := uuid.Must(uuid.FromString(userID))
	dqm := []bob.Mod[*dialect.SelectQuery]{
		sm.Where(psql.Quote("user_id").EQ(psql.Arg(uid))),
		sm.OrderBy(psql.Quote("created_at")).Desc(),
	}
	if limit > 0 {
		dqm = append(dqm, sm.Limit(limit))
	}
	query := models.Tasks.Query(dqm...)
	tt, err := query.All(ctx, db)
	if err != nil {
		xlog.Error("Failed to query tasks", "error", err)
		return nil, err
	}
	return tt, nil
}
