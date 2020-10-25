package main


import (
	"log"
	"net/http"
)





func main() {

	// create file server handler
	fs := http.FileServer( http.Dir( "." ) )

	// start HTTP server with `fs` as the default handler
	server := http.ListenAndServe( "0.0.0.0:9000", fs)
	log.Printf("start http server at: %v", server)


}

