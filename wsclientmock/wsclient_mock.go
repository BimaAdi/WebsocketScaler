package wsclientmock

import "github.com/BimaAdi/WebsocketScaler/core"

type CommandLog struct {
	Command string // send_to_single_user, send_to_multiple_user, send_to_all
	Payload string
}

type MockWSClient struct {
	Logs   []CommandLog
	Event  core.Event
	Scaler core.ScalerContract
}

func NewMockWSClient() *MockWSClient {
	return &MockWSClient{
		Logs:  []CommandLog{},
		Event: nil,
	}
}

func (twsc *MockWSClient) SendToSingleUser(socket_id string, payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_single_user",
		Payload: payload,
	})
}

func (twsc *MockWSClient) SendToMultipleUser(socket_ids []string, payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_multiple_user",
		Payload: payload,
	})
}

func (twsc *MockWSClient) SendToAll(payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_all",
		Payload: payload,
	})
}

func (twsc *MockWSClient) CreateWebsocketRoute(e core.Event, s core.ScalerContract) {
	twsc.Event = e
	twsc.Scaler = s
}

func (twsc *MockWSClient) CallOnConnect(socket_id string, params core.Params) {
	if twsc.Event == nil {
		panic("event not found")
	}

	if twsc.Scaler == nil {
		panic("scaler not found")
	}

	twsc.Event.OnConnect(twsc.Scaler, socket_id, params)
}

func (twsc *MockWSClient) CallOnMessage(socket_id string, payload string) {
	if twsc.Event == nil {
		panic("event not found")
	}

	if twsc.Scaler == nil {
		panic("scaler not found")
	}

	twsc.Event.OnMessage(twsc.Scaler, socket_id, payload)
}

func (twsc *MockWSClient) CallOnDisconnect(socket_id string) {
	if twsc.Event == nil {
		panic("event not found")
	}

	if twsc.Scaler == nil {
		panic("scaler not found")
	}

	twsc.Event.OnDisconnect(twsc.Scaler, socket_id)
}
