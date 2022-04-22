package models

import (
	"time"
)

type OperationRecord struct {
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;"`
	Method       string        `json:"method" form:"method" gorm:"column:method;"`
	Path         string        `json:"path" form:"path" gorm:"column:path;"`
	Status       int           `json:"status" form:"status" gorm:"column:status;"`
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;"`
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;"`
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;"`
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;"`
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;"`
	UserId       uint          `json:"userId" gorm:"column:user_id;default:NULL;"`
	User         User          `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *OperationRecord) TableName() string {
	return "operation_record"
}
