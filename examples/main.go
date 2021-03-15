package main

import (
	`embed`
	`net/http`

	`github.com/flosch/pongo2`
	"github.com/labstack/echo/v4"

	"github.com/mdaliyan/pongo2echo"
)

//go:embed templates/*
var fs embed.FS

func main() {
	e := echo.New()
	e.Renderer = pongo2echo.NewRenderer("templates/", &fs)
	e.Debug = true
	e.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "main.html", pongo2.Context{
			"name": "World",
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
