package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/yakshop/internal/handler"
	yakRepo "github.com/theluckiesthuman/yakshop/internal/repository/implementation"
	yakUseCase "github.com/theluckiesthuman/yakshop/internal/usecase/implementation"
)

func RegisterYakHandlers(g *echo.Group) {
	st := yakRepo.NewYakStore()
	mgr := yakUseCase.NewYakManager(st)
	yh := handler.NewYakHandler(mgr)
	g.POST("/load", yh.Load)
	g.GET("/stock/:T", yh.ViewStock)
	g.GET("/herd/:T", yh.ViewHerd)
	g.POST("/order/:T", yh.Order)
	g.GET("/order-template", yh.OrderTemplate)
}
