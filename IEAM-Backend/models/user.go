package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID          *uint   `gorm:"primary_key; auto_increment; index;" json:"accountId"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	FirstName   *string `json:"firstname"`
	LastName    *string `json:"lastname"`
	PhoneNumber *string `json:"phonenumber"`
	Level       *int    `gorm:"default:0" json:"level"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
