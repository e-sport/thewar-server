package gamenet

import "log"
import (
	"github.com/golang/protobuf/proto"
	"../gameproto"
	"../api"
)

type handler func(msg proto.Message) (error, []byte)

var apiMap map[int]handler = map[int]handler {
	1001: api.HandMethod,
	1002: api.HandMethod,
	1003: api.HandMethod,
	1004: api.HandMethod,
}


func dispatch(id int, msg []byte, conn *connection)  (error){
	pb := &gameproto.Move{}
	err := proto.Unmarshal(msg, pb)
	if err != nil {
		log.Fatal(err)
	}

	f, has := apiMap[id]
	if has {
		if err, b := f(pb); err == nil {
			conn.pushMsg(b);
		}
	}
	return nil
}
