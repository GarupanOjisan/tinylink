package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/garupanojisan/tinylink/pkg/tinylink"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	s := tinylink.NewTinyLinkServer()
	e.GET("/:id", s.RedirectToLongURL)
	e.POST("/shorten", s.CreateNewTinyLink)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
