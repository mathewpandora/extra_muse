package repository

import "extra_muse/internal/model"

type BalanceRepository interface{
	AddBalance(model.NewBalanceAdd) error
}