# WebsocketScaler

Scaling websocket without hassle.

## Table of Contents

**[Installation](#installation)**<br>
**[Quickstart](#quickstart)**<br>
**[Example](#example)**<br>
**[Problem Scaling Websocket](#problem-scaling-websocket)**<br>
**[Core Concept](#core-concept)**<br>
**[Wsclient](#wsclient)**<br>
**[Scaler](#scaler)**<br>
**[Testing](#testing)**<br>

## Installation

```
go get github.com/BimaAdi/WebsocketScaler
```

## Quickstart

For quickstart we gonna use fiber websocket as wsclient and go-redis as scaler. Make sure you have redis installed. For full source code see [QuickstartExample](https://github.com/BimaAdi/WebsocketScaler-example/tree/main/quickstart)

1. Init new go Project

```
mkdir quickstart
cd quickstart
go mod init quickstart
```

2. Install WebsocketScaler

```
go get github.com/BimaAdi/WebsocketScaler
```

3. Install fiber

```
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/contrib/websocket
go get github.com/gofiber/template/html/v2
```

4. Install go-redis

```
go get github.com/redis/go-redis/v9
```

5. Create main.go file

```go
package main

import (
	"context"
	"log"

	"github.com/BimaAdi/WebsocketScaler/core"
	"github.com/BimaAdi/WebsocketScaler/scalergoredis"
	"github.com/BimaAdi/WebsocketScaler/wsclientfiber"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/redis/go-redis/v9"
)

type Event struct {
}

func (e *Event) OnConnect(s core.ScalerContract, socket_id string, params core.Params) {
	s.SendToSingleUser(socket_id, "Welcome")
	s.SendToAll("sommeone connected")
}

func (e *Event) OnMessage(s core.ScalerContract, socket_id string, payload string) {
	s.SendToAll(payload)
}

func (e *Event) OnDisconnect(s core.ScalerContract, socket_id string) {

}

func main() {
    // adjust based on your redis configurations
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	scl := scalergoredis.NewRedisScaler(rdb, ctx, "ws_channel")
	ws_router := wsclientfiber.NewFiberWebsocket()
	go scl.Subscribe(ws_router)
	event := Event{}
	router := ws_router.CreateWebsocketRoute(&event, scl)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/", router)

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}

```

6. Create views/index.html

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>WebsocketScaler Chat</title>
  </head>
  <body>
    <h1>WebsocketScaler Chat</h1>
    <ul id="chats"></ul>
    <input type="text" id="chat_input" />
    <button onclick="sendChat()">send</button>
    <script>
      let chats = [];
      let socket = new WebSocket("ws://localhost:3000/ws/");
      const sendChat = () => {
        let chat_input = document.getElementById("chat_input");
        socket.send(chat_input.value);
        chat_input.value = "";
      };
      const renderChat = (chats) => {
        let ulChats = document.getElementById("chats");
        ulChats.innerHTML = chats.map((x) => "<li>" + x + "</li>").join("");
      };

      socket.onmessage = function (event) {
        chats.push(event.data);
        renderChat(chats);
      };

      socket.onerror = function (error) {
        console.error("[error]");
      };
    </script>
  </body>
</html>
```

7. Run project `go run main.go` open http://localhost:3000 on your browser

## Example

For more example see [WebsocketScalerExample](https://github.com/BimaAdi/WebsocketScaler-example)

## Problem Scaling Websocket

TODO

## Core Concept

TODO

## Wsclient

TODO

## Scaler

TODO

## Testing

TODO
