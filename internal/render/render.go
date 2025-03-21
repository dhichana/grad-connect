package render

import (
    "html/template"
    "io"
    "github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(pattern string) (*TemplateRenderer, error) {
    templates, err := template.ParseGlob(pattern)
    if err != nil {
        return nil, err
    }
    return &TemplateRenderer{templates: templates}, nil
}