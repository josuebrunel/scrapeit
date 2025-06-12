package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"scrapeit/pkg/etherutil"
	"scrapeit/pkg/toast"
	"scrapeit/pkg/view"
	"strings"

	"github.com/a-h/templ"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/josuebrunel/gopkg/xlog"
	"github.com/josuebrunel/gopkg/xsession"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/stephenafamo/bob"
	"github.com/supabase-community/auth-go"
	"github.com/supabase-community/auth-go/types"
)

const (
	// Handler
	HandlerNameDashboard      = "account.dashboard"
	HandlerNameInfo           = "account.info"
	HandlerNamePasswordUpdate = "account.password.update"
	HandlerNameSignIn         = "account.signin"
	HandlerNameSignOut        = "account.signout"
	HandlerNameSignUp         = "account.signup"
	HandlerNameUpdateData     = "account.update.data"

	// Other
	DataKeyEthereumAddress = "ethAddress"
	DataKeyEthereumBalance = "ethBalance"
)

var (
	renderSignUp = func(c echo.Context) templ.Component {
		return SignInAndUpView(
			"Sign Up",
			templ.Attributes{"method": http.MethodPost, "action": view.ReverseX(c, HandlerNameSignUp)},
			true,
		)
	}
	renderSignIn = func(c echo.Context) templ.Component {
		return SignInAndUpView(
			"Sign In",
			templ.Attributes{"method": http.MethodPost, "action": view.ReverseX(c, HandlerNameSignIn)},
			false,
		)
	}
)

type AuthRequest struct {
	Email           string `json:"email" form:"email"`
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"passwordConfirm" form:"passwordConfirm"`
}

type Handler struct {
	name   string
	param  string
	gtc    auth.Client
	db     bob.DB
	ethCli *ethclient.Client
}

func Mount(e *echo.Echo, db *sql.DB, gtc auth.Client, ethCli *ethclient.Client) {
	h := New(db, gtc, ethCli)
	g := e.Group("/account")
	g.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf,header:csrf",
	}))
	g.AddRoute(echo.Route{Method: http.MethodGet, Path: "/signup", Handler: h.SignUp(), Name: HandlerNameSignUp})
	g.AddRoute(echo.Route{Method: http.MethodPost, Path: "/signup", Handler: h.SignUp(), Name: HandlerNameSignUp})
	g.AddRoute(echo.Route{Method: http.MethodGet, Path: "/signin", Handler: h.SignIn(), Name: HandlerNameSignIn})
	g.AddRoute(echo.Route{Method: http.MethodPost, Path: "/signin", Handler: h.SignIn(), Name: HandlerNameSignIn})
	loginRequired := xsession.LoginRequired("/account/signin")
	g.AddRoute(echo.Route{Method: http.MethodGet, Path: "/", Handler: h.Dashboard(), Name: HandlerNameDashboard,
		Middlewares: []echo.MiddlewareFunc{loginRequired}})
	g.AddRoute(echo.Route{Method: http.MethodGet, Path: "/signout", Handler: h.SignOut(), Name: HandlerNameSignOut,
		Middlewares: []echo.MiddlewareFunc{loginRequired}})
	g.AddRoute(echo.Route{Method: http.MethodGet, Path: "/update-data", Handler: h.UpdateData(), Name: HandlerNameUpdateData})
	g.AddRoute(echo.Route{Method: http.MethodPatch, Path: "/update-data", Handler: h.UpdateData(), Name: HandlerNameUpdateData})
}

func New(db *sql.DB, gtc auth.Client, ethCli *ethclient.Client) Handler {
	return Handler{name: "users", param: "uid", gtc: gtc, db: bob.NewDB(db), ethCli: ethCli}
}

func (h Handler) SudoAdmin() auth.Client {
	admin := os.Getenv("APP_ADMIN_EMAIL")
	adminPwd := os.Getenv("APP_ADMIN_PASSWORD")
	user, err := h.gtc.SignInWithEmailPassword(admin, adminPwd)
	if err != nil {
		xlog.Error("failed to authenticate user", "error", err)
		return nil
	}
	return h.gtc.WithToken(user.AccessToken)
}

func (h Handler) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			return view.Render(c, http.StatusOK, renderSignUp(c), nil)
		}
		var req AuthRequest
		if err := c.Bind(&req); err != nil {
			return view.Render(c, http.StatusOK, renderSignUp(c), nil)
		}
		req.PasswordConfirm = req.Password
		_, err := h.gtc.Signup(types.SignupRequest{Email: req.Email, Password: req.Password})
		if err != nil {
			xlog.Error("failed to create user", "error", err)
			authErr, _ := UnmarshalAuthError(err.Error())
			toast.NotifyError(c, authErr.Msg)
			return view.Render(c, http.StatusOK, renderSignUp(c), nil)
		}
		toast.NotifySuccess(c, "Signed up successfully")
		return c.Redirect(http.StatusSeeOther, view.ReverseX(c, HandlerNameSignIn))
	}
}

func (h Handler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			if xsession.IsAuthenticated(c.Request().Context()) {
				toast.NotifyInfo(c, "You are already signed in")
				return c.Redirect(http.StatusSeeOther, view.ReverseX(c, HandlerNameDashboard))
			}
			c.Response().Header().Set("Hx-Redirect", "/account/signin")
			return view.Render(c, http.StatusOK, renderSignIn(c), nil)
		}
		var req AuthRequest
		if err := c.Bind(&req); err != nil {
			return view.Render(c, http.StatusOK, renderSignIn(c), nil)
		}
		resp, err := h.gtc.SignInWithEmailPassword(req.Email, req.Password)
		if err != nil {
			xlog.Debug("failed to authenticate user", "error", err, "response", resp)
			authErr, _ := UnmarshalAuthError(err.Error())
			toast.NotifyError(c, authErr.Msg)
			return view.Render(c, http.StatusOK, renderSignIn(c), nil)
		}
		user := resp.User
		xlog.Debug("user authenticated", "user", user, "token", resp.AccessToken)
		xsession.SetUser(c.Request().Context(), xsession.XUser{
			ID:    user.ID.String(),
			Token: resp.AccessToken, Email: user.Email,
		})
		toast.NotifySuccess(c, "Signed in successfully")
		return c.Redirect(http.StatusSeeOther, view.ReverseX(c, HandlerNameDashboard))
	}
}

func (h Handler) SignOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method != http.MethodGet {
			return view.Render(c, http.StatusOK, renderSignIn(c), nil)
		}
		user := xsession.GetUser(c.Request().Context())
		cli := h.gtc.WithToken(user.Token)
		if err := cli.Logout(); err != nil {
			xlog.Error("failed to logout user", "error", err)
			return c.Redirect(http.StatusSeeOther, view.ReverseX(c, HandlerNameDashboard, user.ID))
		}
		xsession.DeleteUser(c.Request().Context())
		toast.NotifySuccess(c, "Signed out successfully")
		return c.Redirect(http.StatusSeeOther, view.ReverseX(c, HandlerNameSignIn))
	}
}

type RequestUserData struct {
	EthereumAddress string `json:"ethAddress" form:"ethAddress"`
}

func (h Handler) UpdateData() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Request().Context()
			xuser = xsession.GetUser(ctx)
			req   RequestUserData
		)
		if c.Request().Method == http.MethodGet {
			if addr, err := UserEtherAccountGet(ctx, h.db, xuser.ID); err == nil {
				xlog.Debug("user ether account found", "address", addr.Address, "user", xuser.ID)
				xuser.Set(DataKeyEthereumAddress, addr.Address)
				xsession.SetUser(c.Request().Context(), xuser)
			}
			return view.Render(c, http.StatusOK, EthAddressView(xuser, templ.Attributes{}), nil)
		}
		if err := c.Bind(&req); err != nil {
			xlog.Error("failed to bind request", "error", err)
			return view.Render(c, http.StatusOK, EthAddressView(xuser, templ.Attributes{}), nil)
		}
		addr, err := UserEtherAccountUpsert(ctx, h.db, xuser.ID, req.EthereumAddress)
		if err != nil {
			xlog.Error("failed to upsert user ether account", "error", err)
		}
		if addr != nil {
			xuser.Set(DataKeyEthereumAddress, addr.Address)
		}
		xsession.SetUser(c.Request().Context(), xuser)
		return view.Render(c, http.StatusOK, EthAddressView(xuser, templ.Attributes{}), nil)
	}
}

func (h Handler) Dashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Request().Context()
			xuser = xsession.GetUser(ctx)
		)
		if addr, err := UserEtherAccountGet(ctx, h.db, xuser.ID); err == nil {
			xlog.Debug("user ether account found", "address", addr.Address, "user", xuser.ID)
			xuser.Set(DataKeyEthereumAddress, addr.Address)
			balance, _ := etherutil.GetAccountBalance(ctx, h.ethCli, addr.Address)
			xuser.Set(DataKeyEthereumBalance, balance.String())
		}
		xlog.Debug("user authenticated", "user", xuser)
		xsession.SetUser(c.Request().Context(), xuser)
		return view.Render(c, http.StatusOK, DashboardView(xuser, templ.Attributes{}), nil)
	}
}

// AuthErrorResponse defines the structure for the JSON part of the error message.
type AuthErrorResponse struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

// UnmarshalAuthError parses the specific error string format and extracts the AuthErrorResponse.
func UnmarshalAuthError(errorString string) (*AuthErrorResponse, error) {
	// Find the start of the JSON part
	jsonStart := strings.Index(errorString, "{")
	if jsonStart == -1 {
		return nil, fmt.Errorf("could not find start of JSON object in error string: %s", errorString)
	}

	// Extract the JSON substring
	jsonString := errorString[jsonStart:]

	// Define a variable to hold the unmarshalled data
	var authError AuthErrorResponse

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(jsonString), &authError)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from error string: %w", err)
	}

	return &authError, nil
}
