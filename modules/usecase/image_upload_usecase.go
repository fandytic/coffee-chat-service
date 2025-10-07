package usecase

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"coffee-chat-service/modules/model"
)

type ImageUploadUseCase struct{}

func (uc *ImageUploadUseCase) SaveImage(fileHeader *multipart.FileHeader) (*model.ImageUploadResponse, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(fileHeader.Filename))
	filePath := filepath.Join("./public/uploads", fileName)

	if err := os.MkdirAll("./public/uploads", os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file on server: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	imageURL := "/public/uploads/" + fileName

	return &model.ImageUploadResponse{ImageURL: imageURL}, nil
}
