package controllers

import (
	"log"
	"net/http"

	"github.com/filipjedrasik/crr-api/go/models"
	u "github.com/filipjedrasik/crr-api/go/utils"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	// ONLY FOR TESTS
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	//

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	clients[ws] = true

	for {
		var msg models.Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, ws)
		}

		// Save to DB
		userId := r.Context().Value("userId").(uint)
		user := models.GetUser(uint(userId))
		msg.Save(user.ID, user.Name)
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func History(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "Pobrano")
	response["messages"] = models.GetAllMessage()
	u.Respond(w, response)
}
