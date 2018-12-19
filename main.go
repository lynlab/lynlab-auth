package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })

	// GET,POST /apis/v1/**
	e.POST("/apis/v1/register", func(c echo.Context) error {
		var input RegisterInput
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorOutput{http.StatusBadRequest, "Invalid format"})
		}

		e := HandleRegister(&input)
		if e != nil {
			return c.JSON(e.StatusCode, e)
		}
		return c.JSON(http.StatusCreated, "")
	})

	e.POST("/apis/v1/token/generate", func(c echo.Context) error {
		var input TokenGenerateInput
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorOutput{http.StatusBadRequest, "Invalid format"})
		}

		o, e := HandleTokenGenerate(&input)
		if e != nil {
			return c.JSON(e.StatusCode, e)
		}
		return c.JSON(http.StatusOK, o)
	})

	/// HTML static pages.
	// GET /web/**
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.GET("/web/recovery/reset", func(c echo.Context) error {
		if c.QueryParam("recovery_token") == "" {
			return c.String(http.StatusUnauthorized, "Invalid access")
		}
		return c.Render(http.StatusOK, "reset", nil)
	})
	e.GET("/web/recovery/send_email", func(c echo.Context) error { return c.Render(http.StatusOK, "forgot", nil) })
	e.GET("/web/register", func(c echo.Context) error { return c.Render(http.StatusOK, "register", nil) })
	e.GET("/web/signin", func(c echo.Context) error { return c.Render(http.StatusOK, "signin", nil) })

	/// Static files serving.
	e.Static("/statics", "public/statics")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1", "https://auth.lynlab.co.kr"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
