package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
)

type GroupRepository struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) interfaces.GroupRepositoryInterface {
	return &GroupRepository{DB: db}
}

func (r *GroupRepository) CreateGroup(group *entity.ChatGroup) error {
	return r.DB.Create(group).Error
}

func (r *GroupRepository) FindGroupByID(groupID uint) (*entity.ChatGroup, error) {
	var group entity.ChatGroup
	err := r.DB.First(&group, groupID).Error
	return &group, err
}

func (r *GroupRepository) AddMembers(members []entity.ChatGroupMember) error {
	return r.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&members).Error
}

func (r *GroupRepository) GetGroupMembers(groupID uint) ([]entity.ChatGroupMember, error) {
	var members []entity.ChatGroupMember
	err := r.DB.Preload("Customer.Table.Floor").
		Where("chat_group_id = ?", groupID).
		Find(&members).Error
	return members, err
}

func (r *GroupRepository) IsCustomerMember(groupID, customerID uint) (bool, error) {
	var count int64
	err := r.DB.Model(&entity.ChatGroupMember{}).
		Where("chat_group_id = ? AND customer_id = ?", groupID, customerID).
		Count(&count).Error
	return count > 0, err
}

func (r *GroupRepository) CreateGroupMessage(message *entity.GroupChatMessage) error {
	return r.DB.Create(message).Error
}

func (r *GroupRepository) GetGroupMessages(groupID uint, limit int) ([]entity.GroupChatMessage, error) {
	var messages []entity.GroupChatMessage
	err := r.DB.Preload("Sender").
		Where("chat_group_id = ?", groupID).
		Order("created_at desc").
		Limit(limit).
		Find(&messages).Error

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages, err
}

func (r *GroupRepository) FindGroupsByCustomerID(customerID uint) ([]entity.ChatGroupMember, error) {
	var memberships []entity.ChatGroupMember

	err := r.DB.Preload("ChatGroup").
		Where("customer_id = ?", customerID).
		Order("updated_at desc").
		Find(&memberships).Error
	return memberships, err
}
