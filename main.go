package main

import (
	"io"
	"net/http"
	"strings"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func errorView(c echo.Context, status int, message string) error {
	data := map[string]interface{}{
		"status":  status,
		"message": message,
	}
	return c.Render(status, "error", data)
}

func errorAPI(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{"message": message})
}

type AuthedContext struct {
	identity *UserIdentity
	echo.Context
}

func authed(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerTokens := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(bearerTokens) != 2 || bearerTokens[0] != "Bearer" {
			return errorAPI(c, http.StatusUnauthorized, "unauthorized")
		}

		var token UserToken
		db.Where(&UserToken{AccessToken: bearerTokens[1]}).First(&token)
		if token.ID == 0 || token.Expired() {
			return errorAPI(c, http.StatusUnauthorized, "unauthorized")
		}

		return h(AuthedContext{token.GetIdentity(), c})
	}
}

func main() {
	e := echo.New()
	e.Renderer = &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}

	e.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })

	e.POST("/apis/v1/signin", signin)
	e.POST("/apis/v1/register", register)
	e.POST("/apis/v1/authorize", authorize)

	e.GET("/apis/v1/me", authed(getMe))

	e.GET("/web/signin", func(c echo.Context) error {
		errMsg := "Invalid access. Please contact to system manager if this problem persists."
		appID := c.QueryParam("appId")
		redirectURL := c.QueryParam("redirectUrl")
		if appID == "" {
			return errorView(c, http.StatusBadRequest, errMsg)
		}

		var app Application
		db.Where(&Application{UUID: appID}).First(&app)
		if app.ID == 0 || (redirectURL == "" && app.RedirectURL == "") {
			return errorView(c, http.StatusBadRequest, errMsg)
		}

		data := map[string]interface{}{
			"app":         app,
			"redirectURL": redirectURL,
		}
		return c.Render(http.StatusOK, "signin", data)
	})

	e.Static("/statics", "public/statics")

	echo.NotFoundHandler = func(c echo.Context) error {
		return errorView(c, http.StatusNotFound, "Page not found. Please contact to system manager if this problem persists.")
	}
	e.Logger.Fatal(e.Start(":1323"))
}
