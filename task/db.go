package task

import (
	"context"
	"scrapeit/pkg/models"

	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/stephenafamo/bob/dialect/psql/im"
)

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
