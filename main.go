package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	/// HTTP APIs
	// GET /ping
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	/// HTML static pages.
	// GET /web/**
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.GET("/web/forgot", func(c echo.Context) error { return c.Render(http.StatusOK, "forgot", nil) })
	e.GET("/web/register", func(c echo.Context) error { return c.Render(http.StatusOK, "register", nil) })
	e.GET("/web/reset", func(c echo.Context) error { return c.Render(http.StatusOK, "reset", nil) })
	e.GET("/web/signin", func(c echo.Context) error { return c.Render(http.StatusOK, "signin", nil) })

	/// Static files serving.
	e.Static("/statics", "public/statics")

	e.Logger.Fatal(e.Start(":1323"))
}
