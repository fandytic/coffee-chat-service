package websocket

import (
	"encoding/json"
	"log"
	"time"

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
	RecipientID      uint   `json:"recipient_id"`
	Text             string `json:"text"`
	ReplyToMessageID *uint  `json:"reply_to_message_id,omitempty"`
}

type RepliedMessageInfo struct {
	ID         uint   `json:"id"`
	Text       string `json:"text"`
	SenderName string `json:"sender_name"`
}

type IncomingMessagePayload struct {
	MessageID         uint                `json:"message_id"`
	SenderID          uint                `json:"sender_id"`
	SenderName        string              `json:"sender_name"`
	SenderPhotoURL    string              `json:"sender_photo_url"`
	SenderTableNumber string              `json:"sender_table_number"`
	Text              string              `json:"text"`
	Timestamp         time.Time           `json:"timestamp"`
	ReplyTo           *RepliedMessageInfo `json:"reply_to,omitempty"`
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
			// 1. Unmarshal pesan dari klien
			var payload MessagePayload
			if err := json.Unmarshal(directMsg.Message, &payload); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			// 2. Simpan pesan ke database
			chatMessage := entity.ChatMessage{
				SenderID:         directMsg.SenderID,
				RecipientID:      payload.RecipientID,
				Text:             payload.Text,
				ReplyToMessageID: payload.ReplyToMessageID,
			}
			if err := h.DB.Create(&chatMessage).Error; err != nil {
				log.Printf("Failed to save chat message: %v", err)
				continue
			}

			// 3. Ambil detail pengirim untuk payload balasan
			var sender entity.Customer
			if err := h.DB.Preload("Table").First(&sender, directMsg.SenderID).Error; err != nil {
				continue // Abaikan jika pengirim tidak ditemukan
			}

			// 4. (Opsional) Jika ini adalah balasan, ambil info pesan asli
			var repliedToInfo *RepliedMessageInfo
			if chatMessage.ReplyToMessageID != nil {
				var originalMsg entity.ChatMessage
				// Preload("Sender") untuk mendapatkan nama pengirim asli
				if err := h.DB.Preload("Sender").First(&originalMsg, *chatMessage.ReplyToMessageID).Error; err == nil {
					repliedToInfo = &RepliedMessageInfo{
						ID:         originalMsg.ID,
						Text:       originalMsg.Text,
						SenderName: originalMsg.Sender.Name,
					}
				}
			}

			// 5. Cari koneksi penerima
			if recipient, ok := h.customers[payload.RecipientID]; ok {
				// 6. Buat payload lengkap untuk dikirim ke penerima
				responsePayload := IncomingMessagePayload{
					MessageID:         chatMessage.ID,
					SenderID:          sender.ID,
					SenderName:        sender.Name,
					SenderPhotoURL:    sender.PhotoURL,
					SenderTableNumber: sender.Table.TableNumber,
					Text:              chatMessage.Text,
					Timestamp:         chatMessage.CreatedAt,
					ReplyTo:           repliedToInfo,
				}
				responseJSON, _ := json.Marshal(responsePayload)

				// 7. Kirim ke penerima
				recipient.send <- responseJSON
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
