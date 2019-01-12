package actions

import (
	"github.com/gobuffalo/buffalo"
)

var app *buffalo.App

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{})
		app.GET("/", HomeHandler)

		app.GET("/user/index", UserIndex)
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
