package main

import (
	"fmt"
	"golang-websocket-chat-server/db"
	"golang-websocket-chat-server/internal/config"
	"golang-websocket-chat-server/internal/user"
	"golang-websocket-chat-server/internal/ws"
	"golang-websocket-chat-server/router"
	"log"
)

func main() {
	cfg := config.MustLoad()

	dbConnection, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := user.NewRepository(dbConnection.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	if err := router.Start("localhost:8080"); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
