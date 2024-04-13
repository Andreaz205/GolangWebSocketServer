package router

import (
	"github.com/gin-gonic/gin"
	"golang-websocket-chat-server/internal/user"
	"golang-websocket-chat-server/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.POST("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
	r.POST("/ws/storeMessage/:roomId", wsHandler.StoreMessage)
}

func Start(addr string) error {
	return r.Run(addr)
}
