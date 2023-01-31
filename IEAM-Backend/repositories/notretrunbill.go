package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	// "gorm.io/gorm/clause"
	// "gorm.io/gorm/clause"
)



func Getnotreturnbill(id string) ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Raw(`SELECT * FROM forms WHERE account_id = ? AND type_id = 4 AND img_url_bill IS NULL`, id).Scan(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return form, err
	}

	tx.Commit()

	return form, nil
}
