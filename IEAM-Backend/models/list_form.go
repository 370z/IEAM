package models

import (
	"gorm.io/gorm"
	"time"
)

type List_Form struct {
	ID         *uint    `gorm:"primary_key; auto_increment; index;" json:"id_listform"`
	FormId     *int     `json:"-"`
	Title      *string  `json:"title"`
	Unit       *string  `json:"unit"`
	Number     *string  `json:"number"`
	Unit_Price *string  `json:"unit_price"`
	Total      *int `json:"total"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Form *Form `gorm:"foreignkey:FormId" json:"Form,omitempty"`
}
