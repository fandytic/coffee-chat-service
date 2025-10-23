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

func (r *GroupRepository) FindGroupsByCustomerID(customerID uint) ([]entity.ChatGroupMember, error) {
	var memberships []entity.ChatGroupMember

	err := r.DB.Preload("ChatGroup").
		Where("customer_id = ?", customerID).
		Order("updated_at desc").
		Find(&memberships).Error
	return memberships, err
}

func (r *GroupRepository) CountUnreadMessagesPerGroup(customerID uint) (map[uint]int64, error) {
	type UnreadCountResult struct {
		ChatGroupID uint
		UnreadCount int64
	}

	var results []UnreadCountResult

	err := r.DB.Raw(`
		SELECT
			m.chat_group_id,
			COUNT(m.id) as unread_count
		FROM
			group_chat_messages m
		JOIN
			chat_group_members gm ON m.chat_group_id = gm.chat_group_id
		LEFT JOIN
			group_message_read_statuses rs ON m.chat_group_id = rs.chat_group_id AND gm.customer_id = rs.customer_id
		WHERE
			gm.customer_id = ?
			AND m.id > COALESCE(rs.last_read_message_id, 0)
			AND m.sender_id != ?
		GROUP BY
			m.chat_group_id
	`, customerID, customerID).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	unreadMap := make(map[uint]int64)
	for _, result := range results {
		unreadMap[result.ChatGroupID] = result.UnreadCount
	}
	return unreadMap, nil
}

func (r *GroupRepository) MarkGroupMessagesAsRead(customerID, groupID uint) error {
	var lastMessageID uint
	err := r.DB.Model(&entity.GroupChatMessage{}).
		Select("id").
		Where("chat_group_id = ?", groupID).
		Order("id DESC").
		Limit(1).
		Row().Scan(&lastMessageID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	if lastMessageID > 0 {
		readStatus := entity.GroupMessageReadStatus{
			ChatGroupID:       groupID,
			CustomerID:        customerID,
			LastReadMessageID: lastMessageID,
		}

		return r.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "chat_group_id"}, {Name: "customer_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_read_message_id"}),
		}).Create(&readStatus).Error
	}

	return nil
}

func (r *GroupRepository) FindLastGroupMessages(groupIDs []uint) (map[uint]*entity.GroupChatMessage, error) {
	var messages []entity.GroupChatMessage

	if len(groupIDs) == 0 {
		return make(map[uint]*entity.GroupChatMessage), nil
	}

	err := r.DB.Raw(`
		SELECT * FROM group_chat_messages
		WHERE (chat_group_id, id) IN (
			SELECT chat_group_id, MAX(id)
			FROM group_chat_messages
			WHERE chat_group_id IN ?
			GROUP BY chat_group_id
		)
	`, groupIDs).Scan(&messages).Error

	if err != nil {
		return nil, err
	}

	messageMap := make(map[uint]*entity.GroupChatMessage)
	for i, msg := range messages {
		messageMap[msg.ChatGroupID] = &messages[i]
	}

	return messageMap, nil
}
