package scalergochannel_test

import (
	"testing"
	"time"

	"github.com/BimaAdi/WebsocketScaler/scalergochannel"
	"github.com/BimaAdi/WebsocketScaler/wsclientmock"
	"github.com/stretchr/testify/assert"
)

func TestGoChannelScaler(t *testing.T) {
	// Given
	scl := scalergochannel.NewGoChannelScaler()
	ws_router := wsclientmock.NewMockWSClient()
	go scl.Subscribe(ws_router)

	// When
	time.Sleep(2 * time.Second) // wait for start
	scl.SendToSingleUser("a", `{"hello": "a"}`)
	scl.SendToMultipleUser([]string{"a", "b"}, `{"hello": "a and b"}`)
	scl.SendToAll(`{"hello": "all"}`)
	time.Sleep(2 * time.Second) // wait for all message finish

	// Expect
	expect := []wsclientmock.CommandLog{
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
