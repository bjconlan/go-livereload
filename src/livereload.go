package main

import (
	"flag"
	"livereload/http_server"
)

var (
	hostname *string = flag.String("hostname", "localhost", "host name")
	port *uint = flag.Uint("port", 35729, "port number")
)

func main() {
	flag.Parse()
	http_server.Start(*hostname, *port)
}
