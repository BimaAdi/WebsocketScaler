package scaler

import (
	"github.com/BimaAdi/WebsocketScaler"
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

func (gcs GoChannelScaler) Subscribe(ws WebsocketScaler.WSClientContract) {
	isChannelOpen := true
	var v string
	defer close(gcs.c)
	for isChannelOpen {
		v, isChannelOpen = <-gcs.c
		payload, err := WebsocketScaler.UnmarshalMessageToSingleUser(v)
		if err == nil {
			ws.SendToSingleUser(payload.SocketId, payload.Payload)
		}

		payload2, err := WebsocketScaler.UnmarshalMessageToMultipleUser(v)
		if err == nil {
			ws.SendToMultipleUser(payload2.SocketIds, payload2.Payload)
		}

		payload3, err := WebsocketScaler.UnmarshalMessageToaAll(v)
		if err == nil {
			ws.SendToAll(payload3.Payload)
		}
	}
}

func (gcs GoChannelScaler) SendToSingleUser(socket_id string, payload string) {
	data, err := WebsocketScaler.MarshalMessageToSingleUser(socket_id, payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}

func (gcs GoChannelScaler) SendToMultipleUser(socket_ids []string, payload string) {
	data, err := WebsocketScaler.MarshalMessageToMultipleUser(socket_ids, payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}

func (gcs GoChannelScaler) SendToAll(payload string) {
	data, err := WebsocketScaler.MarshalMessageToAll(payload)
	if err != nil {
		panic(err)
	}
	gcs.c <- data
}
