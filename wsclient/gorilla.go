package wsclient

import (
	"log"
	"net/http"

	"github.com/BimaAdi/WebsocketScaler"
	"github.com/gorilla/websocket"
)

type GorillaWebsocket struct {
	Upgrader        websocket.Upgrader
	websocket_conns map[string]*websocket.Conn
}

func NewGorillaWebsocket() GorillaWebsocket {
	return GorillaWebsocket{
		Upgrader:        websocket.Upgrader{},
		websocket_conns: make(map[string]*websocket.Conn),
	}
}

func (gw GorillaWebsocket) SendToSingleUser(socket_id string, payload string) {
	if x, found := gw.websocket_conns[socket_id]; found {
		if err := x.WriteMessage(1, []byte(payload)); err != nil {
			panic(err.Error())
		}
	}
}

func (gw GorillaWebsocket) SendToMultipleUser(socket_ids []string, payload string) {
	for _, socket_id := range socket_ids {
		if x, found := gw.websocket_conns[socket_id]; found {
			if err := x.WriteMessage(1, []byte(payload)); err != nil {
				panic(err.Error())
			}
		}
	}
}

func (gw GorillaWebsocket) SendToAll(payload string) {
	for _, v := range gw.websocket_conns {
		if err := v.WriteMessage(1, []byte(payload)); err != nil {
			panic(err.Error())
		}
	}
}

func (gw *GorillaWebsocket) CreateWebsocketRoute(e WebsocketScaler.Event, s WebsocketScaler.ScalerContract) func(http.ResponseWriter, *http.Request) {
	gw.Upgrader = websocket.Upgrader{}
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := gw.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		socket_id := WebsocketScaler.GenerateRandomString(25)
		gw.websocket_conns[socket_id] = c
		// TODO url Params for gorilla websocket
		// log.Println(r.URL)
		e.OnConnect(s, socket_id, WebsocketScaler.Params{
			Path:        "",
			QueryParams: map[string]string{},
		})
		defer c.Close()

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				delete(gw.websocket_conns, socket_id)
				e.OnDisconnect(s, socket_id)
				break
			}
			e.OnMessage(s, socket_id, string(message))
		}
	}
}
