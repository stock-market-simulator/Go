package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// websocket
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

type Message struct {
	CurrentPrice int `json:"price"`
}

func (a *AppHandler) handleConnections(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print(err)
	}
	defer ws.Close()

	// 클라이언트 등록
	clients[ws] = true

	for {
		var msg Message // db에서 현재가 가져오기

		/*	클라이언트 테스트
			ws = new WebSocket('ws://localhost:5000/ws');
			ws.addEventListener('message', function(e) {
				var msg = JSON.parse(e.data);
				console.log(msg);
			ws.send(JSON.stringify({price:1000}));
		*/
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			return err
		}
		// 가져온 데이터를 브로드캐스트 채널에 보낸다.

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// 브로드캐스트 채널에서 메시지를 기다린다.
		msg := <-broadcast

		// 메시지를 받으면 접속 중인 클라이언트 모두에게 메시지를 보낸다.
		for client := range clients {
			err := client.WriteJSON(msg)

			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
