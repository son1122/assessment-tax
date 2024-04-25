package middleware

import (
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/constant"
	"strings"
)

func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return echo.NewHTTPError(401, "Authentication required")
		}

		authParts := strings.SplitN(authorizationHeader, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Basic" {
			return echo.NewHTTPError(401, "Invalid authentication format")
		}

		encodedCredentials := authParts[1]

		cfg := constant.Get()
		correctCredentials := base64.StdEncoding.EncodeToString([]byte(cfg.AdminUsername + ":" + cfg.AdminPassword))

		if encodedCredentials != correctCredentials {
			return echo.NewHTTPError(401, "Invalid credentials")
		}

		return next(c)
	}
}
