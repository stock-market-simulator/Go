package app

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/stock-market-simulator/Go/db/model"
)

// db 테스트
type AppHandler struct {
	db model.DBHandler
}

func (a *AppHandler) getDbTestHandler(c echo.Context) error {
	user := a.db.GetDbTest()

	return c.JSON(http.StatusOK, user)
}

func (a *AppHandler) createDbTestHandler(c echo.Context) error {
	u := a.db.CreateDbTest()

	return c.JSON(http.StatusOK, u)
}

func MakeHandler() *echo.Echo {
	e := echo.New()
	handler := &AppHandler{
		db: model.NewDBHandler("test"),
	}
	// db 테스트
	e.GET("/test/db/read", handler.getDbTestHandler)
	e.GET("/test/db/create", handler.createDbTestHandler)

	// websocket 테스트
	//e.GET("/test/ws", chatConnections)
	//go handleMessages()

	return e
}

/*
// websocket 테스트
var clients = make(map[*websocket.Conn]bool) // 접속된 클라이언트
var broadcast = make(chan Message)           // 메시지 브로드 캐스트

var upgrader = websocket.Upgrader{} // http 연결을 websocket으로 업그레이드

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func chatConnections(c echo.Context) error {
	// 받은 GET 요청을 websocket으로 업그레이드(ws는 *websockt.Conn 객체임)
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close() // 함수가 끝날 때 websocket을 꼭 닫아줘야함

	// 새로운 클라이언트 등록
	clients[ws] = true

	for {
		var msg Message
		// 새로운 메시지를 JSON으로 읽고, Message 오브젝트에 맵핑
		err := ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
			delete(clients, ws) // clients map에서 오류인 클라이언트 삭제
			break
		}

		// 새롭게 수신된 메시지를 브로드 캐스트 채널에 보낸다.
		broadcast <- msg
	}
	return c.JSON(http.StatusOK, ws)
}

func handleMessages() {
	for {
		// 브로드캐스트 채널에서 다음 메시지를 받는다.(goroutine이므로 계속 메시지가 오는 걸 기다림)
		msg := <-broadcast

		// 현재 접속 중인 클라이언트 모두에게 메시지를 보낸다.
		// map형 이므로 client에는 클라이언트 객체가 할당됨
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client) // clients map에서 오류인 클라이언트 삭제
			}
		}
	}
}
*/
