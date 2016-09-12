package main

import (
	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/api"
	"github.com/toddbluhm/help-people-api/context"
)

func CreateRouter() (router *web.Router) {
	router = web.New(context.Context{}).
		Middleware(web.LoggerMiddleware).             // Use some included middleware
		Middleware(web.ShowErrorsMiddleware).         // ...
		Middleware((*context.Context).SetHelloCount). // Your own middleware!
		Get("/", (*context.Context).SayHello)         // Add a route

	api.AttachRouter(router)

	return
}
