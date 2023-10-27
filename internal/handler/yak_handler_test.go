package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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
		expectedError  error
	}{
		{
			name: "success",
			requestBody: bytes.NewBufferString(`<herd>
			<labyak name="Betty-1" age="4" sex="f" />
			<labyak name="Betty-2" age="8" sex="f" />
			<labyak name="Betty-3" age="9.5" sex="f" />
			</herd>`),
			expectedStatus: http.StatusResetContent,
			expectedError:  nil,
		},
		{
			name:           "error reading request body",
			expectedStatus: http.StatusInternalServerError,
			expectedError:  errors.New("EOF"),
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

			err := yh.Load(c)
			assert.Equal(t, tt.expectedError, err)
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
		expectedError  error
	}{
		{
			name:           "success",
			elapsedDays:    "13",
			expectedStatus: http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "invalid elapsedDays",
			elapsedDays:    "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedError:  errors.New("strconv.Atoi: parsing \"invalid\": invalid syntax"),
		},
		{
			name:           "error viewing stock",
			elapsedDays:    "1000",
			expectedStatus: http.StatusInternalServerError,
			expectedError:  errors.New("no stock available"),
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
