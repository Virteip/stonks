package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIKeyAuth middleware checks for a valid API key in the X-API-Key header
func APIKeyAuth(expectedAPIKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-API-Key")

			if apiKey == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "API key is required",
				})
			}

			if apiKey != expectedAPIKey {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid API key",
				})
			}

			return next(c)
		}
	}
}
