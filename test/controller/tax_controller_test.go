package model_test

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
	"github.com/son1122/assessment-tax/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	e := echo.New()
	router.InitRoutes(e)

	req := httptest.NewRequest(http.MethodPost, "/tax/calculations", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controller.TaxCalculationPost(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
