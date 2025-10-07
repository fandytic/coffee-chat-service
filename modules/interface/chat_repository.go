package interfaces

type ChatRepositoryInterface interface {
	MarkMessagesAsRead(senderID, recipientID uint) error
}
