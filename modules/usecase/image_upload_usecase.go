package usecase

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chai2010/webp"

	"coffee-chat-service/modules/model"
)

type ImageUploadUseCase struct{}

func (uc *ImageUploadUseCase) SaveImage(fileHeader *multipart.FileHeader) (*model.ImageUploadResponse, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	originalFileName := filepath.Base(fileHeader.Filename)
	fileNameWithoutExt := strings.TrimSuffix(originalFileName, filepath.Ext(originalFileName))
	newFileName := fmt.Sprintf("%d_%s.webp", time.Now().Unix(), fileNameWithoutExt)

	uploadDir := "./public/uploads"
	filePath := filepath.Join(uploadDir, newFileName)

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file on server: %w", err)
	}
	defer dst.Close()

	if err := webp.Encode(dst, img, &webp.Options{Quality: 80}); err != nil {
		return nil, fmt.Errorf("failed to encode image to webp: %w", err)
	}

	imageURL := "/public/uploads/" + newFileName

	return &model.ImageUploadResponse{ImageURL: imageURL}, nil
}
