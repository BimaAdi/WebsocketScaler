package scalergoredis_test

import (
	"context"
	"testing"
	"time"

	"github.com/BimaAdi/WebsocketScaler/scalergoredis"
	"github.com/BimaAdi/WebsocketScaler/wsclientmock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestGoRedisScaler(t *testing.T) {
	// Given
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	scl := scalergoredis.NewRedisScaler(rdb, ctx, "ws_channel")
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
