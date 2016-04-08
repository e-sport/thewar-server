package gamenet

import (
	"game/api"
	"game/gameproto"
	"github.com/golang/protobuf/proto"
	"log"
)

type handler func(msg proto.Message) (error, []byte)

var apiMap map[int]handler = map[int]handler{
	1001: api.HandMethod,
	1002: api.HandMethod,
	1003: api.HandMethod,
	1004: api.HandMethod,
}

func getPbBody(id int) proto.Message {
	switch id {
	case 1001:
		return &gameproto.Move{}
	case 1002:
		return &gameproto.Move{}
	case 1003:
		return &gameproto.Move{}
	case 1004:
		return &gameproto.Move{}
	default:
		return &gameproto.Move{}
	}
}

func dispatch(id int, msg []byte, conn *connection) error {
	pb := getPbBody(id)
	err := proto.Unmarshal(msg, pb)
	if err != nil {
		log.Fatal(err)
	}

	f, ok := apiMap[id]
	if ok {
		if err, b := f(pb); err == nil {
			conn.pushMsg(b)
		}
	}
	return nil
}
