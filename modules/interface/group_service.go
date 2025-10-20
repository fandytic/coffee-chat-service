package interfaces

import "coffee-chat-service/modules/model"

type GroupServiceInterface interface {
	CreateGroup(creatorID uint, req model.CreateGroupRequest) (*model.GroupResponse, error)
	InviteMembers(inviterID, groupID uint, req model.InviteToGroupRequest) error
	GetGroupMembers(customerID, groupID uint) ([]model.GroupMemberResponse, error)
	GetCustomerGroups(customerID uint) ([]model.GroupResponse, error)
}
