package app

import (
	"net/http"
	"scrapeit/pkg/models"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/stephenafamo/bob"
)

const RequestContextUserID = "user_id"

var ResponseUnauthorized = Response[string]{Error: "Unauthorized"}

func ApiKeyAuthMiddleware(db bob.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ak := c.Request().Header.Get("Authorization")
			if ak == "" {
				xlog.Debug("No API key provided")
				return c.JSON(http.StatusUnauthorized, ResponseUnauthorized)
			}
			m, err := models.FindAPIKey(c.Request().Context(), db, ak)
			if err != nil {
				xlog.Error("Failed to find API key", "error", err, "api_key", ak)
				return c.JSON(http.StatusUnauthorized, ResponseUnauthorized)
			}
			c.Set(RequestContextUserID, m.UserID.String())
			return next(c)
		}
	}
}
