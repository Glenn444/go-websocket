package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	WebsocketUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct{}

func NewManager() *Manager {
	return  &Manager{}
}


func (m *Manager)serverWs (w http.ResponseWriter, r *http.Request)  {
	var v interface{}
	log.Println("new connection")
	

	//upgrade connection
	conn, err := WebsocketUpgrader.Upgrade(w,r,nil)
	if err != nil{
		fmt.Printf("error occurred in upgrading socket %v",err)
		return
	}
	conn.ReadJSON(v)

	fmt.Printf("data: %v",v)
	conn.Close()
}