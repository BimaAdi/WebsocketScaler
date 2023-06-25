package scaler_test

import (
	"context"
	"testing"
	"time"

	"github.com/BimaAdi/WebsocketScaler/scaler"
	"github.com/BimaAdi/WebsocketScaler/wsclient"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisScaler(t *testing.T) {
	// Given
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	scaler := scaler.NewRedisScaler(rdb, ctx, "ws_channel")
	ws_router := wsclient.NewTestWSClient()
	go scaler.Subscribe(ws_router)

	// When
	time.Sleep(2 * time.Second) // wait for start
	scaler.SendToSingleUser("a", `{"hello": "a"}`)
	scaler.SendToMultipleUser([]string{"a", "b"}, `{"hello": "a and b"}`)
	scaler.SendToAll(`{"hello": "all"}`)
	time.Sleep(2 * time.Second) // wait for all message finish

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
