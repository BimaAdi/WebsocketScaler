package wsclient

import (
	"github.com/BimaAdi/WebsocketScaler/core"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type FiberWebsocket struct {
	websocket_conns map[string]*websocket.Conn
}

func NewFiberWebsocket() FiberWebsocket {
	return FiberWebsocket{
		websocket_conns: make(map[string]*websocket.Conn),
	}
}

func (fw FiberWebsocket) SendToSingleUser(socket_id string, payload string) {
	if x, found := fw.websocket_conns[socket_id]; found {
		if err := x.WriteMessage(1, []byte(payload)); err != nil {
			panic(err.Error())
		}
	}
}

func (fw FiberWebsocket) SendToMultipleUser(socket_ids []string, payload string) {
	for _, socket_id := range socket_ids {
		if x, found := fw.websocket_conns[socket_id]; found {
			if err := x.WriteMessage(1, []byte(payload)); err != nil {
				panic(err.Error())
			}
		}
	}
}

func (fw FiberWebsocket) SendToAll(payload string) {
	for _, v := range fw.websocket_conns {
		if err := v.WriteMessage(1, []byte(payload)); err != nil {
			panic(err.Error())
		}
	}
}

func (fw *FiberWebsocket) CreateWebsocketRoute(e core.Event, s core.ScalerContract) func(c *fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		// 	// c.Locals is added to the *websocket.Conn
		// 	// log.Println(c.Locals("allowed"))  // true
		// 	// log.Println(c.Params("id"))       // 123
		// 	// log.Println(c.Query("v"))         // 1.0
		// 	// log.Println(c.Cookies("session")) // ""
		socket_id := core.GenerateUserId()
		fw.websocket_conns[socket_id] = c
		e.OnConnect(s, socket_id, core.Params{
			Path:        c.Params("id"),
			QueryParams: map[string]string{},
		})

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			// mt  int
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				e.OnDisconnect(s, socket_id)
				delete(fw.websocket_conns, socket_id)
				break
			}

			e.OnMessage(s, socket_id, string(msg))

			// if err = c.WriteMessage(mt, msg); err != nil {
			// 	e.OnDisconnect(socket_id)
			// 	delete(websocket_conns, socket_id)
			// 	break
			// }
		}

	})
}
