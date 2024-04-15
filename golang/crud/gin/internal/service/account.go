package service

import (
	"jixiangup.me/examples/gin/internal/model/dto"
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

func (s AccountService) EditorAccount(payload dto.EditorAccountDTO) (*int64, error) {
	var account entities.Account
	if payload.Id != nil {
		// 编辑账户
		account = entities.Account{Id: payload.Id, Nickname: payload.Nickname, Email: payload.Email}
		err := s.repository.UpdateAccount(account)
		if err != nil {
			return nil, err
		}
	}
	// 创建账户
	account = entities.Account{Nickname: payload.Nickname, Email: payload.Email, Password: payload.Password}
	err := s.repository.InsertAccount(account)
	if err != nil {
		return nil, err
	}
	return account.Id, nil
}
