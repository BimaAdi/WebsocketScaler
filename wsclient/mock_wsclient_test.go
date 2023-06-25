package wsclient_test

import (
	"testing"

	"github.com/BimaAdi/WebsocketScaler"
	"github.com/BimaAdi/WebsocketScaler/scaler"
	"github.com/BimaAdi/WebsocketScaler/wsclient"
	"github.com/stretchr/testify/assert"
)

func TestTestWSClient(t *testing.T) {
	// Given
	ws_router := wsclient.NewMockWSClient()

	// When
	ws_router.SendToSingleUser("a", `{"hello": "a"}`)
	ws_router.SendToMultipleUser([]string{"a", "b"}, `{"hello": "a and b"}`)
	ws_router.SendToAll(`{"hello": "all"}`)

	// Expect
	expect := []wsclient.CommandLog{
		{
			Command: "send_to_single_user",
			Payload: `{"hello": "a"}`,
		},
		{
			Command: "send_to_multiple_user",
			Payload: `{"hello": "a and b"}`,
		},
		{
			Command: "send_to_all",
			Payload: `{"hello": "all"}`,
		},
	}
	for i := 0; i < 3; i++ {
		assert.Equal(t, expect[i].Command, ws_router.Logs[i].Command)
		assert.Equal(t, expect[i].Payload, ws_router.Logs[i].Payload)
	}
}

type Event struct {
}

func (e *Event) OnConnect(s WebsocketScaler.ScalerContract, socket_id string, params WebsocketScaler.Params) {
	s.SendToAll("someone connect with socket_id " + socket_id)

}

func (e *Event) OnMessage(s WebsocketScaler.ScalerContract, socket_id string, payload string) {
	s.SendToAll("got message from " + socket_id + ": " + payload)

}

func (e *Event) OnDisconnect(s WebsocketScaler.ScalerContract, socket_id string) {
	s.SendToAll("user with socket_id " + socket_id + " disconnect")
}

func TestTestWSClientCreateRoute(t *testing.T) {
	// Given
	ws_router := wsclient.NewMockWSClient()
	test_scaler := scaler.NewMockScaler()
	event := Event{}
	test_scaler.Subscribe(ws_router)
	ws_router.CreateWebsocketRoute(&event, test_scaler)

	// When
	ws_router.CallOnConnect("AAAA", WebsocketScaler.Params{})
	ws_router.CallOnMessage("AAAA", "hello everyone")
	ws_router.CallOnDisconnect("AAAA")

	// Expect
	expect := []scaler.CommandLog{
		{
			Command: "send_to_all",
			Payload: "someone connect with socket_id AAAA",
		},
		{
			Command: "send_to_all",
			Payload: "got message from AAAA: hello everyone",
		},
		{
			Command: "send_to_all",
			Payload: "user with socket_id AAAA disconnect",
		},
	}
	for i := 0; i < 3; i++ {
		assert.Equal(t, expect[i].Command, test_scaler.Logs[i].Command)
		assert.Equal(t, expect[i].Payload, test_scaler.Logs[i].Payload)
	}
}
