package main

import (
	"../gameproto"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	//"bytes"
)

func main() {
	url := "ws://localhost:8000/"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	pb := gameproto.Move{Msgid: proto.Int32(1002), X: proto.Int32(1), Y: proto.Int32(2)}
	bin, _ := proto.Marshal(&pb)

	header := make([]byte, 4, 4+len(bin))

	binary.BigEndian.PutUint32(header, 1002)

	/*
		b := new(bytes.Buffer);
		b.Write(header)
		b.Write(bin)
	*/

	c.WriteMessage(websocket.BinaryMessage, append(header, bin...))

}
