package model

import "time"

type DefaultDate struct {
	CreatedAt *time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
	IsDeleted *bool      `json:"is_deleted" gorm:"default:false; not null;"`
}
