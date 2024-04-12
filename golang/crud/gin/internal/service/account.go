package service

import (
	"jixiangup.me/examples/gin/internal/model/entities"
	"jixiangup.me/examples/gin/internal/repository"
)

type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{repository: accountRepository}
}

func (s AccountService) GetAccountDetail(id int64) (*entities.Account, error) {
	return s.repository.GetAccountById(id)
}
