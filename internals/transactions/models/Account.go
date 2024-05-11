package models

import "time"

type Account struct {
	DocId     string
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDeleted bool      `json:"-"`
}
