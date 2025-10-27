package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	WebsocketUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct{
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return  &Manager{
		clients: make(ClientList),
	}
}


func (m *Manager)serverWs (w http.ResponseWriter, r *http.Request)  {
	

	//upgrade connection to websocket
	conn, err := WebsocketUpgrader.Upgrade(w,r,nil)
	if err != nil{
		fmt.Printf("error occurred in upgrading socket %v",err)
		return
	}
	
	client := NewClient(conn,m)
	m.addClient(client)

	//start goroutine to read and write messages
	go client.readMessages()
}


//adds client to the client list
func (m *Manager) addClient(c *Client){
	m.Lock()
	defer m.Unlock()

	m.clients[c] = true
}

//remove client connection if exists in the clients list
func (m *Manager) removeClient(c *Client){
	m.Lock()
	defer m.Unlock()

	if _,ok := m.clients[c]; ok{
		c.connection.Close()
		delete(m.clients,c)
	}
}