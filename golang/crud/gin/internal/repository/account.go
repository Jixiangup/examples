package repository

import (
	"errors"
	"gorm.io/gorm"
	error2 "jixiangup.me/examples/gin/internal/error"
	"jixiangup.me/examples/gin/internal/model/entities"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r AccountRepository) GetAccountById(id int64) (*entities.Account, error) {
	var (
		account entities.Account
		db      = r.db
	)

	result := db.First(&account, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &account, nil
}

func (r AccountRepository) UpdateAccount(account entities.Account) error {
	if &account == nil {
		return error2.NewUnknownMistake("account cannot be nil")
	}
	result := r.db.Model(account).Updates(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r AccountRepository) InsertAccount(account entities.Account) error {
	if &account == nil {
		return error2.NewUnknownMistake("account cannot be nil")
	}
	result := r.db.Create(&account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
