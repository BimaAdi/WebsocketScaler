package wsclient_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/BimaAdi/WebsocketScaler"
	"github.com/BimaAdi/WebsocketScaler/scaler"
	"github.com/BimaAdi/WebsocketScaler/wsclient"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

type EventGorilla struct {
}

func (e *EventGorilla) OnConnect(s WebsocketScaler.ScalerContract, socket_id string, params WebsocketScaler.Params) {
	s.SendToSingleUser(socket_id, "someone connect")
}

func (e *EventGorilla) OnMessage(s WebsocketScaler.ScalerContract, socket_id string, payload string) {
	s.SendToAll("got message " + payload)
}

func (e *EventGorilla) OnDisconnect(s WebsocketScaler.ScalerContract, socket_id string) {
	s.SendToAll("somenone disconnect")
}

func TestGorillaWSClient(t *testing.T) {
	// Given
	ws_router := wsclient.NewGorillaWebsocket()
	test_scaler := scaler.NewMockScaler()
	event := EventGorilla{}
	test_scaler.Subscribe(ws_router)
	router := ws_router.CreateWebsocketRoute(&event, test_scaler)
	s := httptest.NewServer(http.HandlerFunc(router))
	defer s.Close()

	// yank from https://stackoverflow.com/questions/47637308/create-unit-test-for-ws-in-golang
	// When
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()
	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		t.Fatalf("%v", err)
	}
	if err := ws.WriteMessage(websocket.CloseMessage, []byte("disconnect")); err != nil {
		t.Fatalf("%v", err)
	}

	// Expect
	time.Sleep(2 * time.Second) // wait for all message send
	expect := []scaler.CommandLog{
		{
			Command: "send_to_single_user",
			Payload: "someone connect",
		},
		{
			Command: "send_to_all",
			Payload: "got message hello",
		},
		{
			Command: "send_to_all",
			Payload: "somenone disconnect",
		},
	}
	for i := 0; i < 3; i++ {
		assert.Equal(t, expect[i].Command, test_scaler.Logs[i].Command)
		assert.Equal(t, expect[i].Payload, test_scaler.Logs[i].Payload)
	}
}
