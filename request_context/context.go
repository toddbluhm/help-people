package request_context

import (
	"fmt"
	"strings"

	"github.com/gocraft/dbr"
	"github.com/gocraft/web"
)

type ContextMiddleware func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc)

type Context struct {
	HelloCount int
	DB         *dbr.Session
}

func (c *Context) CreateNewDBSessionForRequest(conn *dbr.Connection) web.GenericMiddleware {
	return func(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.DB = conn.NewSession(nil)
		next(rw, req)
	}
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!!!!!!")
}
