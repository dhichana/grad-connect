package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		data := map[string]any{
			"Title":"SSR in Go",
			"Message": "Hello, Server-Side Rendering in Go",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})
}