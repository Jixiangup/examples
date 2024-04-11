package entities

type Account struct {
	Id         *int64          `json:"id" gorm:"primary_key"`
	Nickname   *string         `json:"nickname"`
	Email      *string         `json:"email"`
	Deleted    *bool           `json:"deleted"`
	CreatedAt  *JSONFormatTime `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt *JSONFormatTime `json:"modified_at" gorm:"autoUpdateTime"`
}

func (a Account) TableName() string {
	return "account"
}
