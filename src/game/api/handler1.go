package api

import (
	"game/db"
	"game/gameproto"
	"github.com/golang/protobuf/proto"
	"log"
)

func HandMethod(msg proto.Message) (error, []byte) {
	msg1 := msg.(*gameproto.Move)
	log.Print("HandMethod x=", msg1.GetX())
	db.TestUser()

	return nil, []byte("ok")
}
