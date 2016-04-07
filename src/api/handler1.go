package api

import (
	"github.com/golang/protobuf/proto"
	"log"
	"../gameproto"
)

func HandMethod(msg proto.Message) (error, []byte){
	msg1 := msg.(*gameproto.Move)
	log.Print("HandMethod x=", msg1.GetX())
	return nil, []byte("ok")
}
