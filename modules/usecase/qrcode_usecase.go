package usecase

import (
	"github.com/skip2/go-qrcode"

	"coffee-chat-service/modules/model"
)

type QRCodeUseCase struct{}

func (uc *QRCodeUseCase) GenerateQRCode(req model.QRCodeRequest) ([]byte, error) {
	png, err := qrcode.Encode(req.Content, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
