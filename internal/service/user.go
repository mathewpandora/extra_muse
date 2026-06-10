package service

import (
	"extra_muse/internal/model"
	"extra_muse/internal/repository"
	"fmt"
)

type UserService struct{
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (us *UserService) AddUser (NewUser model.NewUserData) error {

	if err := us.UserRepository.Save(NewUser); err != nil {
		return fmt.Errorf("us.Repo.Save: %w", err)
	}

	return nil
	
}
