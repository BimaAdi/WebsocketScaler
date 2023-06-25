package scaler

import "github.com/BimaAdi/WebsocketScaler"

type CommandLog struct {
	Command string // send_to_single_user, send_to_multiple_user, send_to_all
	Payload string
}

type MockScaler struct {
	Logs []CommandLog
	ws   WebsocketScaler.WSClientContract
}

func NewMockScaler() *MockScaler {
	return &MockScaler{}
}

func (ts *MockScaler) Subscribe(ws WebsocketScaler.WSClientContract) {
	ts.ws = ws
}

func (ts *MockScaler) SendToSingleUser(socket_id string, payload string) {
	ts.ws.SendToSingleUser(socket_id, payload)
	ts.Logs = append(ts.Logs, CommandLog{
		Command: "send_to_single_user",
		Payload: payload,
	})
}

func (ts *MockScaler) SendToMultipleUser(socket_ids []string, payload string) {
	ts.ws.SendToMultipleUser(socket_ids, payload)
	ts.Logs = append(ts.Logs, CommandLog{
		Command: "send_to_multiple_user",
		Payload: payload,
	})
}

func (ts *MockScaler) SendToAll(payload string) {
	ts.ws.SendToAll(payload)
	ts.Logs = append(ts.Logs, CommandLog{
		Command: "send_to_all",
		Payload: payload,
	})
}
