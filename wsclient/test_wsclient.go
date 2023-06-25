package wsclient

type CommandLog struct {
	Command string // send_to_single_user, send_to_multiple_user, send_to_all
	Payload string
}

type TestWSClient struct {
	Logs []CommandLog
}

func NewTestWSClient() *TestWSClient {
	return &TestWSClient{
		Logs: []CommandLog{},
	}
}

func (twsc *TestWSClient) SendToSingleUser(socket_id string, payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_single_user",
		Payload: payload,
	})
}

func (twsc *TestWSClient) SendToMultipleUser(socket_ids []string, payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_multiple_user",
		Payload: payload,
	})
}

func (twsc *TestWSClient) SendToAll(payload string) {
	twsc.Logs = append(twsc.Logs, CommandLog{
		Command: "send_to_all",
		Payload: payload,
	})
}
