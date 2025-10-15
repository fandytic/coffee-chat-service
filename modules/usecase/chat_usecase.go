package usecase

import (
	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type ChatUseCase struct {
	ChatRepo interfaces.ChatRepositoryInterface
}

func (uc *ChatUseCase) MarkMessagesAsRead(senderID, recipientID uint) error {
	return uc.ChatRepo.MarkMessagesAsRead(senderID, recipientID)
}

func (uc *ChatUseCase) GetMessageHistory(user1ID, user2ID uint) ([]model.ChatHistoryMessage, error) {
	messages, err := uc.ChatRepo.GetMessageHistory(user1ID, user2ID)
	if err != nil {
		return nil, err
	}

	history := make([]model.ChatHistoryMessage, 0, len(messages))
	for _, message := range messages {
		history = append(history, buildChatHistoryMessage(&message))
	}

	return history, nil
}

func buildChatHistoryMessage(message *entity.ChatMessage) model.ChatHistoryMessage {
	history := model.ChatHistoryMessage{
		MessageID: message.ID,
		SenderID:  message.SenderID,
		Text:      message.Text,
		Timestamp: message.CreatedAt,
	}

	if message.Sender.ID != 0 {
		history.SenderName = message.Sender.Name
		history.SenderPhotoURL = message.Sender.PhotoURL

		if message.Sender.Table.ID != 0 {
			history.SenderTableNumber = message.Sender.Table.TableNumber

			if message.Sender.Table.Floor.ID != 0 {
				history.SenderFloorNumber = message.Sender.Table.Floor.FloorNumber
			}
		}
	}

	if message.ReplyToMessageID != nil && message.ReplyToMessage != nil {
		repliedSenderName := ""
		if message.ReplyToMessage.Sender.ID != 0 {
			repliedSenderName = message.ReplyToMessage.Sender.Name
		}

		reply := &model.ChatHistoryReply{
			ID:         message.ReplyToMessage.ID,
			Text:       message.ReplyToMessage.Text,
			SenderName: repliedSenderName,
		}

		if message.ReplyToMessage.Menu != nil {
			reply.Menu = &model.ChatHistoryMenu{
				ID:       message.ReplyToMessage.Menu.ID,
				Name:     message.ReplyToMessage.Menu.Name,
				Price:    message.ReplyToMessage.Menu.Price,
				ImageURL: message.ReplyToMessage.Menu.ImageURL,
			}
		}

		history.ReplyTo = reply
	}

	if message.Menu != nil {
		history.Menu = &model.ChatHistoryMenu{
			ID:       message.Menu.ID,
			Name:     message.Menu.Name,
			Price:    message.Menu.Price,
			ImageURL: message.Menu.ImageURL,
		}
	}

	if message.OrderID != nil && message.Order != nil && message.Order.ID != 0 {
		history.Order = buildChatHistoryOrder(message.Order)
	}

	return history
}

func buildChatHistoryOrder(order *entity.Order) *model.ChatHistoryOrder {
	if order == nil {
		return nil
	}

	orderHistory := &model.ChatHistoryOrder{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		NeedType:   order.NeedType,
		TableID:    order.TableID,
		Notes:      order.Notes,
		SubTotal:   order.SubTotal,
		Tax:        order.Tax,
		Total:      order.Total,
	}

	if order.RecipientID != nil {
		orderHistory.RecipientID = order.RecipientID
	}

	if order.Table.ID != 0 {
		orderHistory.TableNumber = order.Table.TableNumber
		orderHistory.TableName = order.Table.TableName
		if order.Table.Floor.ID != 0 {
			orderHistory.TableFloorNumber = order.Table.Floor.FloorNumber
		}
	}

	if len(order.OrderItems) > 0 {
		items := make([]model.ChatHistoryOrderItem, 0, len(order.OrderItems))
		for _, item := range order.OrderItems {
			orderItem := model.ChatHistoryOrderItem{
				ID:       item.ID,
				MenuID:   item.MenuID,
				Quantity: item.Quantity,
				Price:    item.Price,
			}

			if item.Menu.ID != 0 {
				orderItem.Menu = &model.ChatHistoryMenu{
					ID:       item.Menu.ID,
					Name:     item.Menu.Name,
					Price:    item.Menu.Price,
					ImageURL: item.Menu.ImageURL,
				}
			}

			items = append(items, orderItem)
		}
		orderHistory.OrderItems = items
	}

	return orderHistory
}
