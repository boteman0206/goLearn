package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//ws://localhost:8080/ws
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Websocket Connected!")
		listen(websocket)
	})
	http.ListenAndServe(":8080", nil)
}

func listen(conn *websocket.Conn) {

	for {
		// read a message
		messageType, messageContent, err := conn.ReadMessage()
		timeReceive := time.Now()
		fmt.Println("messageType: ", messageType)
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message
		fmt.Println(string(messageContent))

		// reponse message
		messageResponse := fmt.Sprintf("Your message is: %s. Time received : %v", messageContent, timeReceive)

		if err := conn.WriteMessage(messageType, []byte(messageResponse)); err != nil {
			log.Println(err)
			return
		}
	}

}
