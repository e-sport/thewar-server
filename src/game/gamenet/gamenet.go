package gamenet

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8000", "http service address")
var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("senderror: ", err)
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c

	go c.read()
	go c.write()
}

func Start() {
	go h.run()

	http.HandleFunc("/", wsHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("err: ", err)
	}
}
