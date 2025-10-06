package interfaces

import "coffee-chat-service/modules/model"

type QRCodeServiceInterface interface {
	GenerateQRCode(req model.QRCodeRequest) ([]byte, error)
}
