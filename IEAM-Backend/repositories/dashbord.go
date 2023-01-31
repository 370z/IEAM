package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"

	"gorm.io/gorm/clause"
	// "gorm.io/gorm/clause"
)

type dashbord struct {
	Nametype *string `json:"nametype"`
	Sum      *int    `json:"sum"`
}
type dashbord_numberform struct {
	Total *int `json:"total"`
	Approve *int `json:"approve"`
	Cancel *int `json:"cancel"`
	
}

func Getdashbord_numberform() (dashbord_numberform, error) {
	var dashbord_data dashbord_numberform
	tx := config.DB.Begin()

	if err := tx.Debug().Raw(`SELECT count(*) as total,(SELECT count(*) FROM forms WHERE is_approve = 1) as approve,(SELECT count(*) as cancel FROM forms WHERE is_approve = 2)  FROM forms 
	`).Scan(&dashbord_data).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return dashbord_data, err
	}

	tx.Commit()

	return dashbord_data, nil
}
func Getdashbord(id string) ([]dashbord, error) {
	var dashbord_data []dashbord
	tx := config.DB.Begin()

	if err := tx.Debug().Raw(`
	SELECT
	c1.nametype ,c2.sum
FROM
	typeforms c1
	LEFT JOIN (
	SELECT SUM
		( total ),
		forms.account_id,
		type_id 
	FROM
		forms 
	GROUP BY
		forms.type_id,
		forms.account_id, 
		forms.is_approve
	HAVING
	account_id = ? AND is_approve = 1
	) c2 ON c1."id" = c2.type_id ORDER BY "id"`, id).Scan(&dashbord_data).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return dashbord_data, err
	}

	tx.Commit()

	return dashbord_data, nil
}

func Getdashbord_all() ([]dashbord, error) {
	var dashbord_data []dashbord
	tx := config.DB.Begin()

	if err := tx.Debug().Raw(`SELECT
	c1.nametype ,c2.sum
FROM
	typeforms c1
	LEFT JOIN (
	SELECT SUM
		( total ),
		type_id 
	FROM
		forms
	GROUP BY
		forms.type_id,
		forms.is_approve 
	HAVING
	is_approve = 1
	) c2 ON c1."id" = c2.type_id
		ORDER BY "id"`).Scan(&dashbord_data).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return dashbord_data, err
	}

	tx.Commit()

	return dashbord_data, nil
}
func Dashborddata_historynotreturn() ([]models.Form, error) {
	var form []models.Form
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Raw(`SELECT
		A.*
	FROM
		forms A 
	LEFT JOIN forms b ON A.is_return = b.ID 
	WHERE
		A.type_id = 4 
		AND A.is_approve = 1 
		AND b.is_approve IS NULL 
OR b.is_approve != 1`).Find(&form).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return nil, err
	}

	tx.Commit()

	return form, nil
}
