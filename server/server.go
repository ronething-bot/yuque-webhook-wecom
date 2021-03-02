package server

import (
	"yuque-webhook-wecom/config"
	"yuque-webhook-wecom/controller"
	"yuque-webhook-wecom/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wujiyu115/yuqueg"
)

func addMiddleware(e *echo.Echo) {
	// 增加 cors 中间件
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func addApi(e *echo.Echo) {

	//init hook
	hook := controller.NewHook(
		storage.NewYuqueParser(
			storage.NewWeCom(config.Config.GetString("token.wecom")),
			yuqueg.NewService(config.Config.GetString("token.yuque")),
		))
	e.POST("/webhook/:token", hook.Handler) // yuque webhook
}

//CreateEngine echo
func CreateEngine() (*echo.Echo, error) {
	e := echo.New()
	addMiddleware(e)
	addApi(e)

	return e, nil
}
