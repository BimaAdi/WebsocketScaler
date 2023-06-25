package WebsocketScaler

import (
	"encoding/json"
	"errors"
)

type MessageToSingleUser struct {
	MessageToSingleUser bool   `json:"message_to_single_user"`
	SocketId            string `json:"socket_id"`
	Payload             string `json:"payload"`
}

type MessageToMultipleUser struct {
	MessageToMultipleUser bool     `json:"message_to_multiple_user"`
	SocketIds             []string `json:"socket_ids"`
	Payload               string   `json:"payload"`
}

type MessageToAll struct {
	MessageToAll bool   `json:"message_to_all"`
	Payload      string `json:"payload"`
}

func MarshalMessageToSingleUser(socket_id string, payload string) (string, error) {
	data := MessageToSingleUser{
		MessageToSingleUser: true,
		SocketId:            socket_id,
		Payload:             payload,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func MarshalMessageToMultipleUser(socket_ids []string, payload string) (string, error) {
	data := MessageToMultipleUser{
		MessageToMultipleUser: true,
		SocketIds:             socket_ids,
		Payload:               payload,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func MarshalMessageToAll(payload string) (string, error) {
	data := MessageToAll{
		MessageToAll: true,
		Payload:      payload,
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

	if !data.MessageToSingleUser {
		return MessageToSingleUser{}, errors.New("not message to single user json")
	}
	return data, nil
}

func UnmarshalMessageToMultipleUser(message string) (MessageToMultipleUser, error) {
	data := MessageToMultipleUser{}
	err := json.Unmarshal([]byte(message), &data)

	if err != nil {
		return MessageToMultipleUser{}, err
	}

	if !data.MessageToMultipleUser {
		return MessageToMultipleUser{}, errors.New("not message to multiple user json")
	}

	return data, nil
}

func UnmarshalMessageToaAll(message string) (MessageToAll, error) {
	data := MessageToAll{}
	err := json.Unmarshal([]byte(message), &data)

	if err != nil {
		return MessageToAll{}, err
	}

	if !data.MessageToAll {
		return MessageToAll{}, errors.New("not message to all json")
	}

	return data, nil
}
