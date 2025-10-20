package interfaces

type BlockRepositoryInterface interface {
	Block(blockerID, blockedID uint) error
	Unblock(blockerID, blockedID uint) error
	IsBlocked(senderID, recipientID uint) (bool, error)
	GetBlockedList(blockerID uint) (map[uint]bool, error)
}
