package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	DocumentId string    `gorm:"index:doc_id_idx,unique,length:25"`
	ID         uint      `gorm:"primaryKey"`
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

func (a *Account) BeforeSave(tx *gorm.DB) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Account) BeforeUpdate(tx *gorm.DB) error {
	a.UpdatedAt = time.Now()

	return nil
}
