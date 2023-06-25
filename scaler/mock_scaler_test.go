package scaler_test

import (
	"testing"

	"github.com/BimaAdi/WebsocketScaler/scaler"
	"github.com/BimaAdi/WebsocketScaler/wsclient"
	"github.com/stretchr/testify/assert"
)

func TestMockScaler(t *testing.T) {
	// Given
	scl := scaler.NewMockScaler()
	ws_router := wsclient.NewMockWSClient()
	scl.Subscribe(ws_router)

	// When
	scl.SendToSingleUser("a", `{"hello": "a"}`)
	scl.SendToMultipleUser([]string{"a", "b"}, `{"hello": "a and b"}`)
	scl.SendToAll(`{"hello": "all"}`)

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
