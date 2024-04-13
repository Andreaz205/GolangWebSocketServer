package ws

import "fmt"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			fmt.Println("Register new client", cl.RoomID, cl.ID)
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; !ok {
					h.Rooms[cl.RoomID].Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					// broadcast a message saying that the client left the room
					if len(h.Rooms[cl.RoomID].Clients) == 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomId:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomId]; ok {
				fmt.Println(h.Rooms[m.RoomId].Clients)
				for _, cl := range h.Rooms[m.RoomId].Clients {
					fmt.Println(cl)
					cl.Message <- m
				}
			}
		}
	}
}
