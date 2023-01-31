package repositories

import (
	config "ieam-backend/config"
	"ieam-backend/models"
	"time"
)

type ApproveRequest struct {
	ID        *uint `json:"id_form"`
	IsApprove *uint `json:"isapprove"`
}

type TransfermomeyRequest struct {
	ID              *uint   `json:"id_form"`
	IsTransfermomey *bool   `json:"istransfermomey"`
	ImgUrlTransfer  *string `json:"img_urltransfer"`
}

type AddbillRequest struct {
	ID         *uint   `json:"id_form"`
	ImgUrlBill *string `json:"img_urlbill"`
}

func Approve(d ApproveRequest) error {
	tx := config.DB.Begin()

	if err := tx.Debug().Model(&models.Form{}).
		Where("id = ?", d.ID).Updates(map[string]interface{}{
		"is_approve":     d.IsApprove,
		"updatedapprove": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func Transfermomey(d TransfermomeyRequest) error {
	tx := config.DB.Begin()

	if err := tx.Debug().Model(&models.Form{}).
		Where("id = ?", d.ID).Updates(map[string]interface{}{
		"is_transfermomey": d.IsTransfermomey,
		"img_url_transfer": d.ImgUrlTransfer,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
func Addbill(d AddbillRequest) error {
	tx := config.DB.Begin()

	if err := tx.Debug().Model(&models.Form{}).
		Where("id = ?", d.ID).Updates(map[string]interface{}{
		"img_url_bill": d.ImgUrlBill,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
