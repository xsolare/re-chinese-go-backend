package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type OperationRecordService struct{}

/// 																	///

func (operationRecordService *OperationRecordService) CreateOperationRecord(operationRecord models.OperationRecord) (err error) {
	// db := global.GV_DB.Model(&models.OperationRecord{})
	// err = db.Create(operationRecord).Error

	err = global.GV_DB.Create(&operationRecord).Error

	return err
}
