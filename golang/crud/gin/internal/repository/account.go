package repository

import (
	"errors"
	"gorm.io/gorm"
	"jixiangup.me/examples/gin/internal/model/entities"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r AccountRepository) GetUserById(id int64) (*entities.Account, error) {
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
