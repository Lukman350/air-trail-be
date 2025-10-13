package routers

import (
	"air-trail-backend/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var updgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type IWebsocket interface {
	Connect(gin.ResponseWriter, *http.Request, http.Header) error
	Reconnect(gin.ResponseWriter, *http.Request, http.Header) error
	Disconnect()
	SendMessage(any) error
	ReadLoop()
}

type WebSocket struct {
	Name          string
	Connection    *websocket.Conn
	OnReadMessage func(int, []byte, error, *WebSocket)
	BBox          *utils.BBox
}

func (ws *WebSocket) Connect(response gin.ResponseWriter, request *http.Request, responseHeader http.Header) error {
	conn, err := updgrader.Upgrade(response, request, responseHeader)

	if err != nil {
		return err
	}

	ws.Connection = conn

	return nil
}

func (ws *WebSocket) Reconnect(response gin.ResponseWriter, request *http.Request, responseHeader http.Header) error {
	if ws.Connection != nil {
		_ = ws.Connection.Close()
	}
	return ws.Connect(response, request, responseHeader)
}

func (ws *WebSocket) Disconnect() {
	if ws.Connection != nil {
		_ = ws.Connection.Close()
		ws.Connection = nil
	}
	log.Printf("[WS:%s] disconnected", ws.Name)
}

func (ws *WebSocket) SendMessage(msg any) error {
	if ws.Connection == nil {
		return &utils.Cat021Error{Message: fmt.Sprintf("[WS:%s] connection is not open", ws.Name)}
	}
	if err := ws.Connection.WriteJSON(msg); err != nil {
		return &utils.Cat021Error{Message: fmt.Sprintf("[WS:%s] write error: %s", ws.Name, err.Error())}
	}
	return nil
}

func (ws *WebSocket) ReadLoop() {
	if ws.Connection == nil {
		return
	}
	for {
		mt, message, err := ws.Connection.ReadMessage()
		if err != nil {
			ws.OnReadMessage(mt, message, &utils.Cat021Error{Message: fmt.Sprintf("[WS:%s] read error: %s", ws.Name, err.Error())}, ws)
			break
		}
		ws.OnReadMessage(mt, message, nil, ws)
	}
}
