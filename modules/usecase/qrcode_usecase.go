package usecase

import (
	"coffee-chat-service/modules/model"

	"github.com/skip2/go-qrcode"
)

type QRCodeUseCase struct{}

func (uc *QRCodeUseCase) GenerateQRCode(req model.QRCodeRequest) ([]byte, error) {
	// Menghasilkan QR code sebagai gambar PNG dalam bentuk byte slice
	// 256 adalah ukuran gambar (256x256 piksel)
	png, err := qrcode.Encode(req.Content, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
