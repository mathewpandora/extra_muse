package service

import (
	"extra_muse/internal/repository"
	"extra_muse/pkg/polza"
)

type GenerationService struct{
	UserRepository repository.UserRepository
	GenerationRepository repository.GenerationRepository
	PolzaClient polza.PolzaClient
}

func NewGenerataionService(GenRepo repository.GenerationRepository, UserRepo repository.UserRepository, polzaClient polza.PolzaClient) *GenerationService  {
	return &GenerationService{
		UserRepository: UserRepo,
		GenerationRepository: GenRepo,
		PolzaClient: polzaClient,
	}
}



func (gc *GenerationService) Generate() error {

	
 //получаем пользователя из бд
 //получаем его баланс
 //проверяем хватает ли баланса 
}

