package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"scrapeit/pkg/db/models"
	"scrapeit/pkg/view"

	"github.com/aarondl/opt/omit"
	"github.com/gofrs/uuid/v5"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/josuebrunel/gopkg/xsession"
	"github.com/labstack/echo/v5"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/im"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

const numBytes = 32

func generateApiKey() (string, error) {
	bytes := make([]byte, numBytes)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	key := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
	return key, nil
}

func (a App) GenerateApiKey() echo.HandlerFunc {
	return func(c *echo.Context) error {
		var (
			ctx    = c.Request().Context()
			user   = xsession.GetUser(ctx)
			apiKey *models.APIKey
			err    error
		)
		if c.Request().Method == http.MethodPut {
			apiKey, err = apiKeyUpsert(ctx, a.bob, user.ID)
			if err != nil {
				xlog.Error("Error creating API key", "error", err, "user", user.ID)
			}
		} else {
			apiKey, err = apiKeyGetByUserID(ctx, a.bob, user.ID)
			if err != nil {
				xlog.Warn("Error getting API key", "error", err, "user", user.ID)
			}
			if apiKey == nil {
				xlog.Debug("API key not found, creating new one", "user", user.ID)
				apiKey, err = apiKeyUpsert(ctx, a.bob, user.ID)
				if err != nil {
					xlog.Error("Error creating API key", "error", err, "user", user.ID)
				}
			}
		}
		return view.Render(c, http.StatusOK, ApiKeyView(*apiKey), nil)
	}
}

func apiKeyGetByUserID(ctx context.Context, db bob.DB, userID string) (*models.APIKey, error) {
	xlog.Debug("Get API key by user ID", "userID", userID)
	apiKey, err := models.APIKeys.Query(sm.Where(psql.Quote("user_id").EQ(psql.Arg(userID)))).One(ctx, db)
	if err != nil {
		xlog.Error("Error getting API key by user ID", "userID", userID, "error", err)
		return nil, err
	}
	return apiKey, nil
}

func apiKeyUpsert(ctx context.Context, db bob.DB, userID string) (*models.APIKey, error) {
	key, err := generateApiKey()
	if err != nil {
		xlog.Error("Error generating API key", "userID", userID, "error", err)
		return nil, err
	}
	uid, _ := uuid.FromString(userID)
	apiKey, err := models.APIKeys.Insert(&models.APIKeySetter{
		APIKey: omit.From(key),
		UserID: omit.From(uid),
	}, im.OnConflict("user_id").DoUpdate(im.SetExcluded("api_key"))).One(ctx, db)
	if err != nil {
		xlog.Error("Error creating API key", "userID", userID, "error", err)
		return nil, err
	}
	return apiKey, nil
}
