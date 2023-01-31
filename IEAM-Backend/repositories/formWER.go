package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	// "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetallformWer() ([]models.Form, error) {
	var Withdraw_borrow_return []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("type_id IN (3,4,5)").Order("id ASC").Find(&Withdraw_borrow_return).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Withdraw_borrow_return, err
	}

	tx.Commit()

	return Withdraw_borrow_return, nil
}

func GetapproveWBR() ([]models.Form, error) {
	var Withdraw_borrow_return []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("type_id IN (3,4) AND is_approve = 1").Order("id ASC").Find(&Withdraw_borrow_return).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Withdraw_borrow_return, err
	}

	tx.Commit()
	return Withdraw_borrow_return, nil
}