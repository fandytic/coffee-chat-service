package websocket

import (
	"encoding/json"
	"log"

	"coffee-chat-service/modules/entity"

	"gorm.io/gorm"
)

// Message struct untuk komunikasi internal di Hub
type DirectMessage struct {
	SenderID uint
	Message  []byte
}

// Payload JSON yang dikirim antar klien
type MessagePayload struct {
	RecipientID uint   `json:"recipient_id"`
	Text        string `json:"text"`
}
type IncomingMessagePayload struct {
	SenderID          uint   `json:"sender_id"`
	SenderName        string `json:"sender_name"`
	SenderPhotoURL    string `json:"sender_photo_url"`
	SenderTableNumber string `json:"sender_table_number"`
	Text              string `json:"text"`
}
type Hub struct {
	clients    map[*Client]bool
	customers  map[uint]*Client
	incoming   chan *DirectMessage
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	DB         *gorm.DB
}

func NewHub(db *gorm.DB) *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		customers:  make(map[uint]*Client),
		incoming:   make(chan *DirectMessage),
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		DB:         db,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.customers[client.CustomerID] = client
			log.Printf("Client connected: CustomerID %d", client.CustomerID)

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				delete(h.customers, client.CustomerID)
				close(client.send)
				log.Printf("Client disconnected: CustomerID %d", client.CustomerID)
			}

		case directMsg := <-h.incoming:
			var payload MessagePayload
			if err := json.Unmarshal(directMsg.Message, &payload); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			var sender entity.Customer
			err := h.DB.Preload("Table").First(&sender, directMsg.SenderID).Error
			if err != nil {
				log.Printf("Could not find sender with ID %d: %v", directMsg.SenderID, err)
				continue
			}

			if recipient, ok := h.customers[payload.RecipientID]; ok {
				responsePayload := IncomingMessagePayload{
					SenderID:          sender.ID,
					SenderName:        sender.Name,
					SenderPhotoURL:    sender.PhotoURL,
					SenderTableNumber: sender.Table.TableNumber,
					Text:              payload.Text,
				}
				responseJSON, _ := json.Marshal(responsePayload)

				select {
				case recipient.send <- responseJSON:
				default:
					log.Printf("Recipient channel full. Dropping message for CustomerID %d", payload.RecipientID)
				}
			} else {
				log.Printf("Recipient not found or offline: CustomerID %d", payload.RecipientID)
			}

		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
					delete(h.customers, client.CustomerID)
				}
			}
		}
	}
}
