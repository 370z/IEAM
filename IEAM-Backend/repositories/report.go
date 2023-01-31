package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	"gorm.io/gorm/clause"
)

func Getsearch(d models.Search) ([]models.Form, error) {
	var form []models.Form
	

	tx := config.DB.Preload(clause.Associations).Where("is_approve = 1")

	if d.Type != nil {
		tx = tx.Where("type_id = ?", d.Type)
	}
	if d.Month != nil {
		momth := d.Month.Month()
		tx = tx.Where(`EXTRACT ( MONTH FROM updatedapprove ) =  ? `, momth )
	}
	if d.Year != nil {
		year := d.Year.Year()
		tx = tx.Where(`EXTRACT ( YEAR FROM updatedapprove ) = ? `, year)
	}

	if err := tx.Debug().Order("updatedapprove DESC").Find(&form).Error; err != nil {
		return form, err
	}
	return form, nil
}
