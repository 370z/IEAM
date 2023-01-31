package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	// "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Gethistoryuser(id string) ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("account_id = ?", id).Find(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return form, err
	}

	tx.Commit()

	return form, nil
}

func Gethistoryapprove() ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("is_approve IN(1,2)").Find(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return form, err
	}

	tx.Commit()

	return form, nil
}

func Gethistoryfinance() ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("type_id IN(1,2)").Find(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return form, err
	}

	tx.Commit()

	return form, nil
}