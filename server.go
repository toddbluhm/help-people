package main

import "net/http"

func main() {
	router := CreateRouter()
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
