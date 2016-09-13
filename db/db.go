package db

import (
	"fmt"
	"os"

	"github.com/gocraft/dbr"
	"github.com/gocraft/health"
	"github.com/gocraft/web"
	"github.com/toddbluhm/help-people-api/request_context"
)

var connection *dbr.Connection

func Init(stream *health.Stream) (con *dbr.Connection) {
	address := os.Getenv("DB_URL")

	fmt.Println("Opening DB Connection.")
	connection, err := dbr.Open("postgres", address, stream)
	if err != nil {
		panic(err)
	}
	return connection
}

func CreateNewSessionMiddleware(conn *dbr.Connection) request_context.ContextMiddleware {
	return func(c *request_context.Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.DB = conn.NewSession(nil)
		next(rw, req)
	}
}

func Close() {
	fmt.Println("Closing DB Connection.")
	connection.Close()
}
