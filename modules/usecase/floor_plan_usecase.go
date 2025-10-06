package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"

	// "coffee-chat-service/modules/repository"
	"gorm.io/gorm"
)

type FloorPlanUseCase struct {
	DB *gorm.DB
}

// CreateFloorPlan adalah inti logika untuk menyimpan denah dan meja
func (uc *FloorPlanUseCase) CreateFloorPlan(floorNumber int, fileHeader *multipart.FileHeader, tablesDataJSON string) (*model.FloorPlanResponse, error) {
	// 1. Simpan file gambar
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Buat nama file unik untuk menghindari konflik
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(fileHeader.Filename))
	filePath := filepath.Join("./public/uploads", fileName)

	// Pastikan direktori ada
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

	// 2. Unmarshal data meja dari string JSON
	var tablesData []model.TableData
	if err := json.Unmarshal([]byte(tablesDataJSON), &tablesData); err != nil {
		return nil, fmt.Errorf("invalid tables data format: %w", err)
	}

	// 3. Simpan ke database dalam satu transaksi
	var createdFloor entity.Floor
	err = uc.DB.Transaction(func(tx *gorm.DB) error {
		// Buat record floor baru
		floor := entity.Floor{
			FloorNumber: floorNumber,
			ImageURL:    imageURL,
		}
		if err := tx.Create(&floor).Error; err != nil {
			return err
		}

		// Buat record tables yang berelasi dengan floor
		for _, td := range tablesData {
			table := entity.Table{
				TableNumber: td.TableNumber,
				TableName:   td.TableName,
				XCoordinate: td.XCoordinate,
				YCoordinate: td.YCoordinate,
				FloorID:     floor.ID,
			}
			if err := tx.Create(&table).Error; err != nil {
				return err
			}
		}
		createdFloor = floor
		return nil
	})

	if err != nil {
		// Hapus file yang sudah diupload jika transaksi gagal
		os.Remove(filePath)
		return nil, fmt.Errorf("failed to save data to database: %w", err)
	}

	return uc.GetFloorPlanByNumber(createdFloor.FloorNumber)
}

// GetFloorPlanByNumber mengambil data denah dan meja
func (uc *FloorPlanUseCase) GetFloorPlanByNumber(floorNumber int) (*model.FloorPlanResponse, error) {
	var floor entity.Floor
	// Preload memuat data relasi (tables) secara otomatis
	if err := uc.DB.Preload("Tables").First(&floor, "floor_number = ?", floorNumber).Error; err != nil {
		return nil, err
	}

	var tables []model.TableData
	for _, t := range floor.Tables {
		tables = append(tables, model.TableData{
			TableNumber: t.TableNumber,
			TableName:   t.TableName,
			XCoordinate: t.XCoordinate,
			YCoordinate: t.YCoordinate,
		})
	}

	response := &model.FloorPlanResponse{
		ID:          floor.ID,
		FloorNumber: floor.FloorNumber,
		ImageURL:    floor.ImageURL,
		Tables:      tables,
	}

	return response, nil
}

// GetAllFloors mengambil daftar sederhana dari semua lantai yang tersedia.
func (uc *FloorPlanUseCase) GetAllFloors() ([]model.FloorInfoResponse, error) {
	var floors []entity.Floor
	if err := uc.DB.Order("floor_number asc").Find(&floors).Error; err != nil {
		return nil, err
	}

	var response []model.FloorInfoResponse
	for _, floor := range floors {
		response = append(response, model.FloorInfoResponse{
			ID:          floor.ID,
			FloorNumber: floor.FloorNumber,
		})
	}
	return response, nil
}

// UpdateTable memperbarui informasi meja yang ada.
func (uc *FloorPlanUseCase) UpdateTable(tableID uint, req model.UpdateTableRequest) (*entity.Table, error) {
	var table entity.Table
	// Cari meja berdasarkan ID
	if err := uc.DB.First(&table, tableID).Error; err != nil {
		// Jika tidak ditemukan, GORM akan mengembalikan error
		return nil, fmt.Errorf("table with ID %d not found", tableID)
	}

	// Perbarui field-fieldnya
	table.TableName = req.TableName
	table.XCoordinate = req.XCoordinate
	table.YCoordinate = req.YCoordinate

	// Simpan perubahan ke database
	if err := uc.DB.Save(&table).Error; err != nil {
		return nil, err
	}

	return &table, nil
}

// DeleteTable menghapus meja berdasarkan ID (soft delete).
func (uc *FloorPlanUseCase) DeleteTable(tableID uint) error {
	result := uc.DB.Delete(&entity.Table{}, tableID)

	if result.Error != nil {
		return result.Error
	}

	// Cek apakah ada baris yang benar-benar terhapus
	if result.RowsAffected == 0 {
		return fmt.Errorf("table with ID %d not found", tableID)
	}

	return nil
}
