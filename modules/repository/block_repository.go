package repository

import (
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
)

type BlockRepository struct {
	DB *gorm.DB
}

func NewBlockRepository(db *gorm.DB) *BlockRepository {
	return &BlockRepository{DB: db}
}

func (r *BlockRepository) Block(blockerID, blockedID uint) error {
	block := entity.CustomerBlock{
		BlockerID: blockerID,
		BlockedID: blockedID,
	}
	return r.DB.FirstOrCreate(&block, block).Error
}

func (r *BlockRepository) Unblock(blockerID, blockedID uint) error {
	return r.DB.Unscoped().Where("blocker_id = ? AND blocked_id = ?", blockerID, blockedID).Delete(&entity.CustomerBlock{}).Error
}

func (r *BlockRepository) IsBlocked(senderID, recipientID uint) (bool, error) {
	var count int64
	err := r.DB.Model(&entity.CustomerBlock{}).
		Where("blocker_id = ? AND blocked_id = ?", recipientID, senderID).
		Count(&count).Error
	return count > 0, err
}

func (r *BlockRepository) GetBlockedList(blockerID uint) (map[uint]bool, error) {
	var blockedUsers []entity.CustomerBlock
	if err := r.DB.Where("blocker_id = ?", blockerID).Find(&blockedUsers).Error; err != nil {
		return nil, err
	}
	blockedMap := make(map[uint]bool)
	for _, user := range blockedUsers {
		blockedMap[user.BlockedID] = true
	}
	return blockedMap, nil
}
