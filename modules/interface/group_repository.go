package interfaces

import "coffee-chat-service/modules/entity"

type GroupRepositoryInterface interface {
	CreateGroup(group *entity.ChatGroup) error
	FindGroupByID(groupID uint) (*entity.ChatGroup, error)
	AddMembers(members []entity.ChatGroupMember) error
	GetGroupMembers(groupID uint) ([]entity.ChatGroupMember, error)
	IsCustomerMember(groupID, customerID uint) (bool, error)
	CreateGroupMessage(message *entity.GroupChatMessage) error
	FindGroupsByCustomerID(customerID uint) ([]entity.ChatGroupMember, error)
}
