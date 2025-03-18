package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()


	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("../web/templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"Title":   "SSR in Go",
			"Message": "Hello, Server-Side Rendering in Go!",
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
