package main

import (
	"flag"
	"livereload/http_server"
)

var (
	hostname = flag.String("hostname", "", "host name")
	port = flag.Uint("port", 35729, "port number")
	directories = Flag.Args()
)

func main() {
	flag.Parse()

	if len(directories) == 0 {
		directories = make([]string, 1, 1)
		directories[1] = "."
	}

	http_server.Start(*hostname, *port)
}
