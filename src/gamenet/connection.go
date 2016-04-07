package gamenet

import (
	"log"
	"github.com/gorilla/websocket"
	"encoding/binary"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)


type connection struct {
	ws *websocket.Conn
	send chan []byte
}

func (c *connection) pushMsg(msg []byte){
	c.send <- msg;
}

func (c *connection) read(){
	defer func() {
		h.unregister <- c
		c.ws.Close()
	}()

	for {
		mt, message, err := c.ws.ReadMessage()

		if len(message) > 4 {
			id := binary.BigEndian.Uint32(message[:4])
			payload := message[4:]
			dispatch( int(id), payload, c);
		}

		if err != nil {
			log.Println("read:", mt, err)
			break
		}
		//log.Printf("recv: %s", message)
	}
}

func (c *connection) write(){
	ticker := time.NewTicker(pongWait)
	defer func() {
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.ws.WriteMessage(websocket.PongMessage, []byte{}); err != nil {
				log.Print(err)
				return
			}
		}
	}
}
