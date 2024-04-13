package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomId   string `json:"roomId"`
	Username string `json:"username"`
}

func (c *Client) writeMessage() {
	defer func() {
		if err := c.Conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			fmt.Println("not ok ws message")
			return
		}

		if err := c.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomId:   c.RoomID,
			Username: c.Username,
		}

		fmt.Println(msg)

		hub.Broadcast <- msg
	}
}
