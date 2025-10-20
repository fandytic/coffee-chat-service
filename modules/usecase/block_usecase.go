package usecase

import (
	interfaces "coffee-chat-service/modules/interface"
)

type BlockUseCase struct {
	BlockRepo interfaces.BlockRepositoryInterface
}

func (uc *BlockUseCase) BlockCustomer(blockerID, blockedID uint) error {
	return uc.BlockRepo.Block(blockerID, blockedID)
}

func (uc *BlockUseCase) UnblockCustomer(blockerID, blockedID uint) error {
	return uc.BlockRepo.Unblock(blockerID, blockedID)
}
