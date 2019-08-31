package server

import (
	"net/http"
	"strconv"

	base "github.com/gloryof/go-openapi-practice/module/base/api"
	"github.com/gloryof/go-openapi-practice/module/user/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server サーバを表す構造体
type Server struct {
	// echo echoの構造体
	echo *echo.Echo
	// port 起動ポート
	port int
}

// NewServer サーバを作成する
func NewServer() Server {

	s := Server{
		echo: echo.New(),
		port: 8000,
	}

	s.middleware()
	s.route()

	return s
}

// Start サーバの起動を行う
func (s *Server) Start() {

	s.echo.Start(":" + strconv.Itoa(s.port))
}

// middleware middlewareの設定
func (s *Server) middleware() {

	s.echo.Use(middleware.KeyAuthWithConfig(authKeyConfig()))
	s.echo.Use(checkCommonHeader())
}

// route APIルートの設定
func (s *Server) route() {

	s.echo.GET("/users", api.ListUsers)
	s.echo.POST("/users", api.RegisterUser)

	s.echo.GET("/users/:id", api.GetUser)
	s.echo.PUT("/users/:id", api.UpdatetUser)
	s.echo.DELETE("/users/:id", api.DeletetUser)
}

// authKeyConfig 認証キーの設定
func authKeyConfig() middleware.KeyAuthConfig {

	return middleware.KeyAuthConfig{
		Validator: getValidator(),
	}
}

// getValidator KeyAuthValidatorを取得する
func getValidator() middleware.KeyAuthValidator {
	return func(key string, c echo.Context) (bool, error) {
		return key == "test", nil
	}
}

// checkCommonHeader 共通ヘッダのチェック処理
func checkCommonHeader() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			h := c.Request().Header
			v := h.Get("X-API-VERSION")

			if v == "" {
				return c.JSON(http.StatusBadRequest, base.ErrorResponse{
					Summary: "入力データに不備があります",
					Details: []string{
						"X-API-VERSIONが設定されていません。",
					},
				})
			}

			return next(c)
		}
	}
}
