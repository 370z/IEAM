package models

import (
	"gorm.io/gorm"
	"time"
)

type Typeform struct {
	ID          *uint   `gorm:"primary_key; auto_increment; index;" json:"typeId"`
	Nametype    *string `json:"nametype"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}


