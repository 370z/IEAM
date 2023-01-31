package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	// "gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetallformIe() ([]models.Form, error) {
	var income_expenses_all []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("type_id IN (1,2)").Order("id ASC").Find(&income_expenses_all).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return income_expenses_all, err
	}

	tx.Commit()

	return income_expenses_all, nil
}