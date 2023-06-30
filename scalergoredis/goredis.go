package scalergoredis

import (
	"context"

	"github.com/BimaAdi/WebsocketScaler/core"
	"github.com/redis/go-redis/v9"
)

type RedisScaler struct {
	Rdb     *redis.Client
	Ctx     context.Context
	Channel string
}

func NewRedisScaler(rdb *redis.Client, ctx context.Context, channel string) RedisScaler {
	return RedisScaler{
		Rdb:     rdb,
		Ctx:     ctx,
		Channel: channel,
	}
}

func (rs RedisScaler) Subscribe(ws core.WSClientContract) {
	pubsub := rs.Rdb.Subscribe(rs.Ctx, rs.Channel)
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(rs.Ctx)
		if err != nil {
			panic(err)
		}

		payload, err := core.UnmarshalMessageToSingleUser(msg.Payload)
		if err == nil {
			ws.SendToSingleUser(payload.SocketId, payload.Payload)
		}

		payload2, err := core.UnmarshalMessageToMultipleUser(msg.Payload)
		if err == nil {
			ws.SendToMultipleUser(payload2.SocketIds, payload2.Payload)
		}

		payload3, err := core.UnmarshalMessageToaAll(msg.Payload)
		if err == nil {
			ws.SendToAll(payload3.Payload)
		}
	}

}

func (rs RedisScaler) SendToSingleUser(socket_id string, payload string) {
	data, err := core.MarshalMessageToSingleUser(socket_id, payload)
	if err != nil {
		panic(err)
	}
	rs.Rdb.Publish(rs.Ctx, rs.Channel, data)
}

func (rs RedisScaler) SendToMultipleUser(socket_ids []string, payload string) {
	data, err := core.MarshalMessageToMultipleUser(socket_ids, payload)
	if err != nil {
		panic(err)
	}
	rs.Rdb.Publish(rs.Ctx, rs.Channel, data)
}

func (rs RedisScaler) SendToAll(payload string) {
	data, err := core.MarshalMessageToAll(payload)
	if err != nil {
		panic(err)
	}
	rs.Rdb.Publish(rs.Ctx, rs.Channel, data)
}
