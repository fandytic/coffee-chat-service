package interfaces

type BlockServiceInterface interface {
	BlockCustomer(blockerID, blockedID uint) error
	UnblockCustomer(blockerID, blockedID uint) error
}
