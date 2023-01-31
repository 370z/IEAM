package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"
	"ieam-backend/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Getallform() ([]models.Form, error) {
	var income_expenses_all []models.Form
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("is_approve = 0").Order("created_at DESC").Find(&income_expenses_all).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return income_expenses_all, err
	}

	tx.Commit()

	return income_expenses_all, nil
}

func Getdetailform(id string) ([]models.Form, error) {
	var income_expenses_all []models.Form
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Where("id = ?", id).Find(&income_expenses_all).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return income_expenses_all, err
	}

	tx.Commit()

	return income_expenses_all, nil
}

func Createform(d models.Form) (models.Form, error) {
	tx := config.DB.Begin()
	radom := utils.RandomString(10) 
	err := Getnumber_from(radom);
	for err != nil {
		radom := utils.RandomString(10)
		err = Getnumber_from(radom)
	}
	d.Number_form = &radom
	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}
func Createformreturn(d models.Form) (models.Form, error) {
	tx := config.DB.Begin()
	radom := utils.RandomString(10) 
	err := Getnumber_from(radom);
	for err != nil {
		radom := utils.RandomString(10)
		err = Getnumber_from(radom)
	}
	d.Returnborrow.Number_form = &radom
	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func Updateform(d models.Form) (models.Form, error) {
	tx := config.DB.Begin()

	if len(*d.List) > 0 {
		for _, list := range *d.List {
			if err := tx.Debug().Where("id != ?  AND form_id = ?", list.ID, d.ID).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.List_Form{}).Error; err != nil {
				tx.Rollback()
			}
		}
	} else {
		if err := tx.Debug().Where("form_id = ?", d.ID).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.List_Form{}).Error; err != nil {
			tx.Rollback()
		}
	}

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func Deleteform(id string) error {

	tx := config.DB.Begin()

	if err := tx.Debug().Where("id = ? ", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Form{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Debug().Where("form_id = ?", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.List_Form{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func Getnumber_from(number_form string) error {
	var income_expenses_all []models.Form
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Where("number_form = ?", number_form).Find(&income_expenses_all).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return err
	}

	tx.Commit()

	return nil
}
