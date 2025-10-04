package websocket

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// writePump memompa pesan dari hub ke koneksi websocket.
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		message, ok := <-c.send
		if !ok {
			// Hub menutup channel.
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error writing message: %v", err)
			return
		}
	}
}

// readPump memompa pesan dari koneksi websocket ke hub.
// Fungsi ini hanya untuk menjaga koneksi dan menangani pemutusan.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		if _, _, err := c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

// ServeWs menangani permintaan websocket dari peer.
// Fungsi ini dipanggil dari middleware Fiber di main.go.
func ServeWs(hub *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
		client.hub.register <- client

		go client.writePump()
		client.readPump() // Jalankan readPump di goroutine utama untuk menjaga koneksi
	}
}
