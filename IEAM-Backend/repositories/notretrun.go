package repositories

import (
	"gorm.io/gorm/clause"
	config "ieam-backend/config"
	"ieam-backend/models"
)

func Getnotreturn(id string) ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Raw(`SELECT
		A.*
	FROM
		forms A 
	LEFT JOIN forms b ON A.is_return = b.ID 
	WHERE
		A.account_id = ?
		AND A.type_id = 4 
		AND A.is_approve = 1 
		AND b.is_approve IS NULL 
OR b.is_approve != 1`, id).Find(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return form, err
	}

	tx.Commit()

	return form, nil
}