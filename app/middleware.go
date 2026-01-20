package app

import (
	"net/http"
	"scrapeit/pkg/db/models"

	"github.com/josuebrunel/gopkg/xlog"
	"github.com/labstack/echo/v5"
	"github.com/stephenafamo/bob"
)

const (
	RequestContextUserID  = "user_id"
	RequestContextEthAddr = "eth_address"
)

var (
	ResponseUnauthorized = Response[string]{Error: "Unauthorized"}
	ResponseNoEthAddr    = Response[string]{Error: "No Ethereum address set"}
)

func ApiKeyAuthMiddleware(db bob.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			var (
				ctx = c.Request().Context()
				ak  = c.Request().Header.Get("Authorization")
			)
			if ak == "" {
				xlog.Debug("No API key provided")
				return c.JSON(http.StatusUnauthorized, ResponseUnauthorized)
			}
			m, err := models.FindAPIKey(ctx, db, ak)
			if err != nil {
				xlog.Error("Failed to find API key", "error", err, "api_key", ak)
				return c.JSON(http.StatusUnauthorized, ResponseUnauthorized)
			}
			c.Set(RequestContextUserID, m.UserID.String())
			ethAddr, err := models.FindEtherAccount(ctx, db, m.UserID)
			if err != nil {
				xlog.Error("Failed to find user ether account", "error", err)
				return c.JSON(http.StatusUnauthorized, ResponseNoEthAddr)
			}
			c.Set(RequestContextEthAddr, ethAddr.Address)
			return next(c)
		}
	}
}
