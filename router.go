package main

import (
	"github.com/gocraft/dbr"
	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api"
	"github.com/toddbluhm/help-people-api/db"
	"github.com/toddbluhm/help-people-api/request_context"
)

func CreateRouter(database *dbr.Connection) (router *web.Router) {
	router = web.New(request_context.Context{}).
		Middleware(web.LoggerMiddleware).     // Use some included middleware
		Middleware(web.ShowErrorsMiddleware). // ...
		Middleware(db.CreateNewSessionMiddleware(database)).
		Middleware((*request_context.Context).SetHelloCount). // Your own middleware!
		Get("/", (*request_context.Context).SayHello)         // Add a route

	api.AttachRouter(router)

	return
}
