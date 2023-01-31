package models

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	ID        *uint `gorm:"primary_key; auto_increment; index;" json:"historyId"`
	FormId    *uint `json:"-"`
	AccountId *uint `gorm:"not null; index;" json:"-"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Form *Form `gorm:"foreignkey:FormId" json:"form,omitempty"`
	// User *User `gorm:"foreignkey:AccountId" json:"user,omitempty"`
}
