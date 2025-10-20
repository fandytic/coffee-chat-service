package usecase

import (
	"encoding/json"
	"time"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/websocket"
)

type BellUseCase struct {
	CustomerRepo interfaces.CustomerRepositoryInterface
	Hub          *websocket.Hub
}

func (uc *BellUseCase) CallWaiter(customerID uint) error {
	customer, err := uc.CustomerRepo.FindCustomerWithDetails(customerID)
	if err != nil {
		return err
	}

	bellPayload := model.BellNotificationPayload{
		CustomerName: customer.Name,
		TableNumber:  customer.Table.TableNumber,
		FloorNumber:  customer.Table.Floor.FloorNumber,
		CallTime:     time.Now(),
	}

	notification := map[string]interface{}{
		"type": "WAITER_BELL",
		"data": bellPayload,
	}
	notificationJSON, _ := json.Marshal(notification)

	uc.Hub.BroadcastToAdmins(notificationJSON)

	return nil
}
