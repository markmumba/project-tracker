package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WebsocketController struct {
	CommunicationService services.CommunicationService
	clients              map[uint]*websocket.Conn
}

func NewWebsocketController(commService services.CommunicationService) *WebsocketController {
	return &WebsocketController{
		CommunicationService: commService,
		clients:              make(map[uint]*websocket.Conn),
	}
}

func (wsc *WebsocketController) HandleWebsockets(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	userID, err := helpers.ConvertUserID(c, "userId")

	wsc.clients[userID] = ws

	for {
		var msg models.CommunicationHistory
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(wsc.clients, userID)
			break
		}
		err = wsc.CommunicationService.SaveMessage(&msg)
		if err != nil {
			log.Printf("error saving message : %v", err)
		}

		if client, ok := wsc.clients[msg.ReceiverID]; ok {
			err = client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(wsc.clients, msg.ReceiverID)
			}
		}
	}

	return nil
}
