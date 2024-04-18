package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/model"
	"net/http"
)

// GetUsers handles the GET users route
// @Summary Show account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} string "ok"
// @Router /accounts/{id} [get]
func GetUsers(c echo.Context) error {
	// Assume a getUsers function fetches users and returns them
	users, err := model.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}
	return c.JSON(http.StatusOK, users)
}
