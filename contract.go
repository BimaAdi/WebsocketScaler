package WebsocketScaler

type Params struct {
	Path        string
	QueryParams map[string]string
}

type Event interface {
	OnConnect(s ScalerContract, socket_id string, param Params)
	OnMessage(s ScalerContract, socket_id string, payload string)
	OnDisconnect(s ScalerContract, socket_id string)
}

type ScalerContract interface {
	SendToSingleUser(socket_id string, payload string)
	SendToMultipleUser(socket_ids []string, payload string)
	SendToAll(payload string)
}

type WSClientContract interface {
	SendToSingleUser(socket_id string, payload string)
	SendToMultipleUser(socket_ids []string, payload string)
	SendToAll(payload string)
}
