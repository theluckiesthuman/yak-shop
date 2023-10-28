package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CloudyKit/jet/v6"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/entities"
	ys "github.com/theluckiesthuman/yakshop/internal/repository/implementation"
	"github.com/theluckiesthuman/yakshop/internal/usecase/implementation"
	"github.com/theluckiesthuman/yakshop/internal/util/dump"
)

func TestYakHandlerLoad(t *testing.T) {
	store := ys.NewYakStore()
	mgr := implementation.NewYakManager(store)
	tests := []struct {
		name           string
		requestBody    io.Reader
		expectedStatus int
	}{
		{
			name: "success",
			requestBody: bytes.NewBufferString(`<herd>
			<labyak name="Betty-1" age="4" sex="f" />
			<labyak name="Betty-2" age="8" sex="f" />
			<labyak name="Betty-3" age="9.5" sex="f" />
			</herd>`),
			expectedStatus: http.StatusResetContent,
		},
		{
			name:           "error reading request body",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", tt.requestBody)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			yh := &YakHandler{
				ym: mgr,
			}

			_ = yh.Load(c)
			cupaloy.SnapshotT(t, dump.HttpResponseDump(rec.Result()))

		})
	}
}

func TestYakHandlerViewStock(t *testing.T) {
	store := ys.NewYakStore()
	herd := entities.Herd{
		Yaks: []entities.Yak{
			{
				Name: "Betty 1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty 2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty 8",
				Age:  9.5,
				Sex:  "f",
			},
		},
	}
	store.Store(herd)
	mgr := implementation.NewYakManager(store)
	yh := &YakHandler{
		ym: mgr,
	}

	tests := []struct {
		name           string
		elapsedDays    string
		expectedStatus int
	}{
		{
			name:           "success",
			elapsedDays:    "13",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid elapsedDays",
			elapsedDays:    "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error viewing stock",
			elapsedDays:    "1000",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("T")
			c.SetParamValues(tt.elapsedDays)

			_ = yh.ViewStock(c)
			cupaloy.SnapshotT(t, dump.HttpResponseDump(rec.Result()))

		})
	}
}

func TestYakHandlerViewHerd(t *testing.T) {
	store := ys.NewYakStore()
	herd := entities.Herd{
		Yaks: []entities.Yak{
			{
				Name: "Betty 1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty 2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty 8",
				Age:  9.5,
				Sex:  "f",
			},
		},
	}
	store.Store(herd)
	mgr := implementation.NewYakManager(store)
	yh := &YakHandler{
		ym: mgr,
	}

	tests := []struct {
		name           string
		elapsedDays    string
		expectedStatus int
	}{
		{
			name:           "success",
			elapsedDays:    "13",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid elapsedDays",
			elapsedDays:    "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error viewing herd",
			elapsedDays:    "1000",
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("T")
			c.SetParamValues(tt.elapsedDays)

			_ = yh.ViewHerd(c)

			cupaloy.SnapshotT(t, dump.HttpResponseDump(rec.Result()))
		})
	}
}

func TestYakHandlerOrder(t *testing.T) {
	store := ys.NewYakStore()
	herd := entities.Herd{
		Yaks: []entities.Yak{
			{
				Name: "Betty 1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty 2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty 8",
				Age:  9.5,
				Sex:  "f",
			},
		},
	}
	store.Store(herd)
	mgr := implementation.NewYakManager(store)
	yh := &YakHandler{
		ym: mgr,
	}

	tests := []struct {
		name           string
		elapsedDays    string
		customerOrder  dto.CustomerOrder
		expectedStatus int
	}{
		{
			name:        "success",
			elapsedDays: "14",
			customerOrder: dto.CustomerOrder{
				Customer: "Medvedev",
				Order: dto.Order{
					Milk:  1100,
					Skins: 3,
				},
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:        "partial order",
			elapsedDays: "14",
			customerOrder: dto.CustomerOrder{
				Customer: "Medvedev",
				Order: dto.Order{
					Milk:  1200,
					Skins: 3,
				},
			},
			expectedStatus: http.StatusPartialContent,
		},
		{
			name:        "no stock available",
			elapsedDays: "10",
			customerOrder: dto.CustomerOrder{
				Customer: "Medvedev",
				Order: dto.Order{
					Milk:  12000,
					Skins: 333,
				},
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:        "invalid elapsedDays",
			elapsedDays: "invalid",
			customerOrder: dto.CustomerOrder{
				Customer: "Medvedev",
				Order: dto.Order{
					Milk:  12000,
					Skins: 333,
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			b, err := json.Marshal(tt.customerOrder)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("T")
			c.SetParamValues(tt.elapsedDays)

			err = yh.Order(c)
			cupaloy.SnapshotT(t, dump.HttpResponseDump(rec.Result()))
		})
	}
}

func loadTemplate(path string) {
	views = jet.NewSet(
		jet.NewOSFileSystemLoader(path),
		jet.InDevelopmentMode(),
	)
}

func TestYakHandlerOrderTemplate(t *testing.T) {

	tests := []struct {
		name         string
		templatePath string
	}{
		{
			name:         "success",
			templatePath: "../../html",
		},
		{
			name: "error getting template",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadTemplate(tt.templatePath)
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			yh := &YakHandler{}

			_ = yh.OrderTemplate(c)
			cupaloy.SnapshotT(t, dump.HttpResponseDump(rec.Result()))
		})
	}
}
