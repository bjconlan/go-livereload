package main

import (
	"flag"
	"fmt"
	"gopkg.in/fsnotify.v0"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)
const (
	API_VERSION = "1.6"
)
var (
	hostname = flag.String("hostname", "", "host name")
	port = flag.Uint("port", 35729, "port number")
	directories = flag.Args()

	// http channels
	subscriptionChannel = make(chan subscriptionMessage)
	messageChannel = make(chan string)
	watcher, watcher_err = fsnotify.NewWatcher()
)

type subscriptionMessage struct {
	websocket *websocket.Conn
	active bool
	url string
}

type livereloadClientUpdateMessage struct {
	path string
	apply_js_live bool
	apply_css_live bool
	apply_images_live bool
}

func eventHub() {
	connections := make(map[*websocket.Conn]string)
	for {
		select {
			case subscriptionMessage := <-subscriptionChannel:
				if subscriptionMessage.active { // on connection
					connections[subscriptionMessage.websocket] = subscriptionMessage.url
				} else { // on disconnection
					delete(connections, subscriptionMessage.websocket)
				}
			case message := <-messageChannel:
				if len(message) == 0 {
					fmt.Printf("close ws requested")
				}
				fmt.Printf("Incomming Message: %s", message)
			case ev := <-watcher.Event:
				//update := livereloadClientUpdateMessage{ev.Name, true, true, true}
				for k, _ :=  range connections {
					websocket.Message.Send(k, fmt.Sprintf("[\"refresh\": {\"path\": \"%s\"}]", ev.Name))
					//websocket.JSON.Send(k, update)
				}
				/*'["refresh", {
					"path" : %s,
					"apply_js_live": true,
					"apply_css_live" : true
					"apply_images_live" : true
				}]'*/
				log.Println("fsevent:", ev)
			case err := <-watcher.Error:
				log.Println("fserror:", err)
		}
	}
}

func liveReload16ConnectionHandler(ws *websocket.Conn) {
	defer func() {
		subscriptionChannel <- subscriptionMessage{ws, false, ""}
		fmt.Printf("Browser disconnected")
		ws.Close()
	}()

	websocket.Message.Send(ws, fmt.Sprintf("!!ver:%s", API_VERSION))
	fmt.Printf("Browser Connected")

	// on connection it's the client url
	var onConnectionMessage string
	websocket.Message.Receive(ws, &onConnectionMessage)
	subscriptionChannel <- subscriptionMessage{ws, true, onConnectionMessage}
	fmt.Printf("Browser URL: %s", onConnectionMessage)

	// websocket messages from the clients get pushed though the eventhub
	for {
		var msg string
		websocket.Message.Receive(ws, &msg)
		messageChannel <- msg
	}
}

func websocketHandshakeHandler(w http.ResponseWriter, req *http.Request) {
	websocket.Handler(liveReload16ConnectionHandler).ServeHTTP(w, req)
}

func main() {
	if watcher_err != nil {
		log.Fatal("Unable to create file monitor.")
	}

	flag.Parse()

	// register directories to monitory
	if len(directories) == 0 {
		directories = make([]string, 1, 1)
		directories[0] = "."
	}

	for _, directory := range directories {
		err := watcher.Watch(directory)
		if err != nil {
			log.Fatal(err)
		}
	}

	go eventHub()
	http.HandleFunc("/websocket", websocketHandshakeHandler)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *hostname, *port), nil)
	if err != nil {
		log.Fatal("http_server.Start.ListenAndServe: ", err)
	}
}
