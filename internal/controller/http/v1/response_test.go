package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	responseErrorJson   = `{"message":"bar"}`
	responseSuccessJson = `{"data":{"foo":"bar"}}`
)

func Test_responseError(t *testing.T) {
	loggerBckp := loggerFunc
	defer func() {
		loggerFunc = loggerBckp
	}()
	logEmitted := false
	loggerFunc = func(err error, statusCode int, msg string) {
		logEmitted = true
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, responseError(c, http.StatusInternalServerError, errors.New("foo"), "bar")) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, responseErrorJson, strings.TrimRight(rec.Body.String(), "\n"))
		assert.Equal(t, true, logEmitted)
	}
}

func Test_responseSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, responseSuccess(c, http.StatusOK, map[string]interface{}{"foo": "bar"})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseSuccessJson, strings.TrimRight(rec.Body.String(), "\n"))
	}
}
