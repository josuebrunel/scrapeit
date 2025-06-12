package account

import (
	"context"
	"scrapeit/pkg/models"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/im"
)

func UserEtherAccountGet(ctx context.Context, db bob.DB, userID string) (*models.EtherAccount, error) {
	addr, err := models.FindEtherAccount(ctx, db, uuid.Must(uuid.FromString(userID)))
	if err != nil {
		xlog.Error("Failed to find user ether account", "error", err)
		return &models.EtherAccount{}, err
	}
	return addr, nil
}

func UserEtherAccountUpsert(ctx context.Context, db bob.DB, userID, address string) (*models.EtherAccount, error) {
	uid := uuid.Must(uuid.FromString(userID))
	addr, err := models.EtherAccounts.Insert(&models.EtherAccountSetter{
		UserID:  &uid,
		Address: &address,
	}, im.OnConflict("user_id").DoUpdate(
		im.SetExcluded("address"),
		im.SetCol("updated_at").ToArg(time.Now().UTC()),
	)).One(ctx, db)
	if err != nil {
		xlog.Error("Failed to upsert user ether account", "error", err)
		return &models.EtherAccount{}, err
	}
	return addr, nil
}
