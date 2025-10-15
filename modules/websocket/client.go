package websocket

import (
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	send       chan []byte
	CustomerID uint // Akan 0 jika ini adalah admin
	AdminID    uint // Akan 0 jika ini adalah customer
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message, ok := <-c.send
		if !ok {
			return
		}

		c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error writing message: %v", err)
			return
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.incoming <- &DirectMessage{SenderID: c.CustomerID, Message: message}
	}
}

func ServeWs(hub *Hub, customerID uint) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), CustomerID: customerID}
		client.hub.register <- client

		go client.writePump()
		client.readPump()
	}
}

func ServeCustomerWs(hub *Hub, customerID uint) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), CustomerID: customerID}
		client.hub.register <- client
		go client.writePump()
		client.readPump()
	}
}

func ServeAdminWs(hub *Hub, adminID uint) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), AdminID: adminID}
		client.hub.register <- client
		go client.writePump()

		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				client.hub.unregister <- client
				break
			}
		}
	}
}
