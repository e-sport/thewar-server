package api

import (
	"../gameproto"
	"github.com/golang/protobuf/proto"
	"log"
)

func HandMethod(msg proto.Message) (error, []byte) {
	msg1 := msg.(*gameproto.Move)
	log.Print("HandMethod x=", msg1.GetX())
	return nil, []byte("ok")
}
