package usecase

import (
	"errors"
	"log"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type GroupUseCase struct {
	GroupRepo interfaces.GroupRepositoryInterface
}

func NewGroupUseCase(groupRepo interfaces.GroupRepositoryInterface) interfaces.GroupServiceInterface {
	return &GroupUseCase{GroupRepo: groupRepo}
}

func (uc *GroupUseCase) CreateGroup(creatorID uint, req model.CreateGroupRequest) (*model.GroupResponse, error) {
	if req.Name == "" {
		return nil, errors.New("group name is required")
	}

	group := &entity.ChatGroup{
		Name:      req.Name,
		CreatorID: creatorID,
	}
	if err := uc.GroupRepo.CreateGroup(group); err != nil {
		return nil, err
	}

	memberIDs := append(req.MemberIDs, creatorID)
	var members []entity.ChatGroupMember
	uniqueIDs := make(map[uint]bool)

	for _, id := range memberIDs {
		if !uniqueIDs[id] {
			members = append(members, entity.ChatGroupMember{
				ChatGroupID: group.ID,
				CustomerID:  id,
			})
			uniqueIDs[id] = true
		}
	}

	if err := uc.GroupRepo.AddMembers(members); err != nil {
		return nil, errors.New("failed to add initial members")
	}

	return &model.GroupResponse{
		ID:        group.ID,
		Name:      group.Name,
		CreatorID: group.CreatorID,
	}, nil
}

func (uc *GroupUseCase) InviteMembers(inviterID, groupID uint, req model.InviteToGroupRequest) error {
	isMember, err := uc.GroupRepo.IsCustomerMember(groupID, inviterID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("forbidden: you are not a member of this group")
	}

	var newMembers []entity.ChatGroupMember
	for _, id := range req.CustomerIDs {
		newMembers = append(newMembers, entity.ChatGroupMember{
			ChatGroupID: groupID,
			CustomerID:  id,
		})
	}

	if len(newMembers) > 0 {
		return uc.GroupRepo.AddMembers(newMembers)
	}
	return nil
}

func (uc *GroupUseCase) GetGroupMembers(customerID, groupID uint) ([]model.GroupMemberResponse, error) {
	isMember, err := uc.GroupRepo.IsCustomerMember(groupID, customerID)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("forbidden: you are not a member of this group")
	}

	members, err := uc.GroupRepo.GetGroupMembers(groupID)
	if err != nil {
		return nil, err
	}

	response := make([]model.GroupMemberResponse, 0, len(members))
	for _, member := range members {
		if member.Customer.ID != 0 {
			response = append(response, model.GroupMemberResponse{
				CustomerID:  member.Customer.ID,
				Name:        member.Customer.Name,
				PhotoURL:    member.Customer.PhotoURL,
				TableNumber: member.Customer.Table.TableNumber,
				FloorNumber: member.Customer.Table.Floor.FloorNumber,
			})
		}
	}

	return response, nil
}

func (uc *GroupUseCase) GetCustomerGroups(customerID uint) ([]model.GroupResponse, error) {
	memberships, err := uc.GroupRepo.FindGroupsByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	unreadMap, err := uc.GroupRepo.CountUnreadMessagesPerGroup(customerID)
	if err != nil {
		log.Printf("Warning: could not retrieve group unread counts: %v", err)
		unreadMap = make(map[uint]int64)
	}

	response := make([]model.GroupResponse, 0, len(memberships))
	for _, member := range memberships {
		if member.ChatGroup.ID != 0 {
			response = append(response, model.GroupResponse{
				ID:          member.ChatGroup.ID,
				Name:        member.ChatGroup.Name,
				CreatorID:   member.ChatGroup.CreatorID,
				UnreadCount: unreadMap[member.ChatGroup.ID],
			})
		}
	}
	return response, nil
}
