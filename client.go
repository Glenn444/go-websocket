package main

import (
	"log"

	"github.com/gorilla/websocket"
)


type ClientList map[*Client]bool


type Client struct{
	connection *websocket.Conn
	manager *Manager
}

func NewClient(conn *websocket.Conn,m *Manager)*Client{
	return &Client{
		connection: conn,
		manager: m,
	}
}

func (c *Client) readMessages(){
	defer func()  {
		//cleanup connection when there is an error
		c.manager.removeClient(c)
	}()

	for{
		messageType, p, errr := c.connection.ReadMessage()

		if errr != nil{
			log.Printf("error occured now %v\n",errr)
			break
		}
		log.Print(messageType)
		log.Print(string(p))
		log.Printf("clientList: %v",c.manager.clients)
	}
}