package main

import (
	_ "database/sql"
	"net/http"
	"os"

	"github.com/gocraft/health"
	_ "github.com/lib/pq"
	"github.com/toddbluhm/help-people-api/db"
)

// Save the stream as a global variable
var stream = health.NewStream()

func main() {
	// Log to stdout! (can also use WriterSink to write to a log file, Syslog, etc)
	stream.AddSink(&health.WriterSink{os.Stdout})

	dbConnection := db.Init(stream)
	defer dbConnection.Close()

	router := CreateRouter(dbConnection)
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
