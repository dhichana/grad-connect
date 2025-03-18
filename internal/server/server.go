package server

import (
	"github.com/dhichana/grad-connect/internal/handlers"
	"github.com/dhichana/grad-connect/internal/render"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) error{
	e := echo.New()
	renderer, err := render.NewTemplateRenderer("../web/templates/*.html")
    if err != nil {
        return err
    }
    e.Renderer = renderer

	handlers.RegisterRoutes(e, db)

    return e.Start(":5000")
}