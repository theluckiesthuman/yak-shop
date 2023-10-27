package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/mapper"
	"github.com/theluckiesthuman/yakshop/internal/usecase/contract"
)

var (
	views = jet.NewSet(
		jet.NewOSFileSystemLoader("./html"),
		jet.InDevelopmentMode(),
	)
)

type YakHandler struct {
	ym contract.YakManager
}

func NewYakHandler(ym contract.YakManager) *YakHandler {
	return &YakHandler{ym}
}

func (yh *YakHandler) Load(ctx echo.Context) error {
	req := ctx.Request()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	herd, err := mapper.MapReqBodyToHerd(body)
	if err != nil {
		log.Println("Error mapping XML:", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	yh.ym.Store(*herd)
	ctx.Response().WriteHeader(http.StatusResetContent)
	return nil
}

func (yh *YakHandler) ViewStock(ctx echo.Context) error {
	elapsedDays := ctx.Param("T")
	elapsedDaysInt, err := strconv.Atoi(elapsedDays)
	if err != nil {
		log.Println("Invalid elapsedDays", err)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	st, err := yh.ym.ViewStock(elapsedDaysInt)
	if err != nil {
		log.Println("Error viewing stock:", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, st)
}

func (yh *YakHandler) ViewHerd(ctx echo.Context) error {
	elapsedDays := ctx.Param("T")
	elapsedDaysInt, err := strconv.Atoi(elapsedDays)
	if err != nil {
		log.Println("Invalid elapsedDays", err)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	hd, err := yh.ym.ViewHerd(elapsedDaysInt)
	if err != nil {
		log.Println("Error viewing herd:", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, hd)
}

func (yh *YakHandler) Order(ctx echo.Context) error {
	days := ctx.Param("T")
	var co dto.CustomerOrder
	if err := ctx.Bind(&co); err != nil {
		log.Println("Error binding request:", err)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	daysInt, err := strconv.Atoi(days)
	if err != nil {
		log.Println("Invalid days", err)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return err
	}
	or, status := yh.ym.Order(daysInt, co)
	if status == dto.Fulfilled {
		ctx.JSON(http.StatusCreated, or)
	} else if status == dto.Partial {
		ctx.JSON(http.StatusPartialContent, or)
	} else {
		//failed to fulfill order
		ctx.JSON(http.StatusNoContent, or)
	}
	return nil
}

func (yh *YakHandler) OrderTemplate(ctx echo.Context) error {
	view, err := views.GetTemplate("order.jet")
	if err != nil {
		log.Println("Error getting template:", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	err = view.Execute(ctx.Response().Writer, nil, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}
