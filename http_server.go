package main

import (
	"log"
	"net/http"
)

func main() {

	// create file server handler
	fs := http.FileServer( http.Dir( "http_server_files" ) )

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe( ":9000", fs ))

}