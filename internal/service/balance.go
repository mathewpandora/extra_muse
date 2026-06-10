package service

import (
	"extra_muse/internal/model"
	"extra_muse/internal/repository"
	"fmt"
)

type BalanceService struct{
	BalanceRepository repository.BalanceRepository
}

func NewBalanceService(repo repository.BalanceRepository) *BalanceService {
	return &BalanceService{
		BalanceRepository: repo,
	}
}


func (bc *BalanceService) AddBalance (amount model.NewBalanceAdd) error {
	if err := bc.BalanceRepository.AddBalance(amount); err != nil {
		return fmt.Errorf("service.AddBalance: %w", err)
	}

	return nil
}