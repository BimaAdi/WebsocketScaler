package WebsocketScaler

import "encoding/json"

type MessageToSingleUser struct {
	SocketId string `json:"socket_id"`
	Payload  string `json:"payload"`
}

type MessageToMultipleUser struct {
	SocketIds []string `json:"socket_ids"`
	Payload   string   `json:"payload"`
}

type MessageToAll struct {
	Payload string `json:"payload"`
}

func MarshalMessageToSingleUser(socket_id string, payload string) (string, error) {
	data := MessageToSingleUser{
		SocketId: socket_id,
		Payload:  payload,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func MarshalMessageToMultipleUser(socket_ids []string, payload string) (string, error) {
	data := MessageToMultipleUser{
		SocketIds: socket_ids,
		Payload:   payload,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func MarshalMessageToAll(payload string) (string, error) {
	data := MessageToAll{
		Payload: payload,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func UnmarshalMessageToSingleUser(message string) (MessageToSingleUser, error) {
	data := MessageToSingleUser{}
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		return MessageToSingleUser{}, err
	}
	return data, nil
}

func UnmarshalMessageToMultipleUser(message string) (MessageToMultipleUser, error) {
	data := MessageToMultipleUser{}
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		return MessageToMultipleUser{}, err
	}
	return data, nil
}

func UnmarshalMessageToaAll(message string) (MessageToAll, error) {
	data := MessageToAll{}
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		return MessageToAll{}, err
	}
	return data, nil
}
