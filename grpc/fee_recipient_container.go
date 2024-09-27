package grpc

import (
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

type FeeRecipientContainer struct {
	mu               *sync.RWMutex
	nextFeeRecipient common.Address
}

func NewFeeRecipientContainer() FeeRecipientContainer {
	return FeeRecipientContainer{
		mu:               &sync.RWMutex{},
		nextFeeRecipient: common.Address{},
	}
}

func (frc *FeeRecipientContainer) GetNextFeeRecipient() common.Address {
	frc.mu.RLock()
	defer frc.mu.RUnlock()
	return frc.nextFeeRecipient
}

func (frc *FeeRecipientContainer) SetNextFeeRecipient(nextFeeRecipient common.Address) {
	frc.mu.Lock()
	defer frc.mu.Unlock()
	frc.nextFeeRecipient = nextFeeRecipient
}
