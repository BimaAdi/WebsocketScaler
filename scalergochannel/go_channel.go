package scalergochannel

import (
	"github.com/BimaAdi/WebsocketScaler/core"
)

type GoChannelScaler struct {
	c chan string
}

func NewGoChannelScaler() GoChannelScaler {
	c := make(chan string)
	return GoChannelScaler{
		c: c,
	}
}

func (gcs GoChannelScaler) Subscribe(ws core.WSClientContract) {
	isChannelOpen := true
	var v string
	defer close(gcs.c)
	for isChannelOpen {
		v, isChannelOpen = <-gcs.c
		payload, err := core.UnmarshalMessageToSingleUser(v)
		if err == nil {
			ws.SendToSingleUser(payload.SocketId, payload.Payload)
		}

		payload2, err := core.UnmarshalMessageToMultipleUser(v)
		if err == nil {
			ws.SendToMultipleUser(payload2.SocketIds, payload2.Payload)
		}

		payload3, err := core.UnmarshalMessageToaAll(v)
		if err == nil {
			ws.SendToAll(payload3.Payload)
		}
	}
}

func (gcs GoChannelScaler) SendToSingleUser(socket_id string, payload string) {
	data, err := core.MarshalMessageToSingleUser(socket_id, payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}

func (gcs GoChannelScaler) SendToMultipleUser(socket_ids []string, payload string) {
	data, err := core.MarshalMessageToMultipleUser(socket_ids, payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}

func (gcs GoChannelScaler) SendToAll(payload string) {
	data, err := core.MarshalMessageToAll(payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}
