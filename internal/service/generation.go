package service

import (
	"errors"
	"extra_muse/internal/model"
	"extra_muse/internal/repository"
	"extra_muse/pkg/polza"
)

type GenerationService struct {
	UserRepository       repository.UserRepository
	GenerationRepository repository.GenerationRepository
	PolzaClient          polza.PolzaClient
	PricePerGen          float64
}

func NewGenerataionService(GenRepo repository.GenerationRepository, UserRepo repository.UserRepository, polzaClient polza.PolzaClient, PricePerGen float64) *GenerationService {
	return &GenerationService{
		UserRepository:       UserRepo,
		GenerationRepository: GenRepo,
		PolzaClient:          polzaClient,
		PricePerGen:          PricePerGen,
	}
}

func (gc *GenerationService) Generate(GenData model.NewGenerationData) error {
	user, err := gc.UserRepository.GetById(GenData.TgID)

	if err != nil {
		return err
	}

	if user.Balance < float64(gc.PricePerGen) {
		return errors.New("You dont have enoufgh money to pay")
	}

}
