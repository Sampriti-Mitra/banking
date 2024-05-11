package models

import (
	"errors"
	"time"
)

type Account struct {
	DocumentId string    `gorm:"index:doc_id_idx,unique,length:25"`
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	IsDeleted  bool      `json:"-"`
}

func (a *Account) Validate() error {
	if a.DocumentId == "" {
		return errors.New("document id cannot be empty")
	}
	return nil
}
