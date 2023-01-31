package models

import (
	"gorm.io/gorm"
	"time"
)

type Form struct {
	ID              *uint      `gorm:"primary_key; auto_increment; index;" json:"id_form"`
	Title           *string    `json:"title"`
	Number_form     *string    `json:"number_form"`
	AccountId       *uint      `gorm:"not null; index;" json:"userId"`
	TypeId          *uint      `gorm:"not null; index;" json:"typeId"`
	Total           *int       `json:"total"`
	IsApprove       *uint      `gorm:"default:0" json:"isapprove"`
	IsReturn        *uint      `gorm:"index;" json:"isreturn"`
	IsTransfermomey *bool      `gorm:"default:false" json:"istransfermoney"`
	ImgUrl          *string    `json:"img_url"`
	ImgUrlTransfer  *string    `json:"img_urltransfer"`
	ImgUrlBill      *string    `json:"img_urlbill"`
	Updatedapprove  *time.Time `gorm:"DEFAULT:now()" json:"updatedapprove"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"CreatedAt"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Typeform     *Typeform    `gorm:"foreignkey:TypeId" json:"type,omitempty"`
	Account      *Account     `gorm:"foreignkey:AccountId" json:"user"`
	List         *[]List_Form `gorm:"foreignkey:FormId" json:"list,omitempty"`
	Returnborrow *Form        `gorm:"foreignkey:IsReturn" json:"returnborrow,omitempty"`
}
