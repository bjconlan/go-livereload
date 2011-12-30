package http_server

import (
	"io"
	"strings"
	"strconv"
	"http"
	"websocket"
	"log"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "testing yall!\n")
}

func LiveReloadServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func Start(hostname string, port uint) {
	http.HandleFunc("/", RootHandler)
	http.Handle("/livereload", websocket.Handler(LiveReloadServer))
	err := http.ListenAndServe(
			strings.Join([]string{hostname, ":", strconv.Uitoa(port)}, ""), nil)
	if err != nil {
		log.Fatal("http_server.Start.ListenAndServe: ", err.String())
	}
}
