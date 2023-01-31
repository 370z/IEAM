package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Getalluser() ([]models.Account, error) {
	var user []models.Account
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Find(&user).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return user, err
	}

	tx.Commit()
	return user, nil
}

func GetuserMe(userId float64) (models.Account, error) {
	var user models.Account
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("id= ?", userId).Find(&user).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return user, err
	}

	tx.Commit()

	return user, nil
}

func Getuserdetail(username string) (models.Account, error) {
	var user models.Account
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("username = ?", username).Find(&user).Error; err != nil {
		tx.Commit()
		return user, err
	}
	tx.Commit()

	return user, nil
}

func Createuser(d models.Account) (models.Account, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func Editeuser(d models.Account) (models.Account, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func DeleteUser(id string) error {

	tx := config.DB.Begin()

	if err := tx.Debug().Where("id = ? ", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Account{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}