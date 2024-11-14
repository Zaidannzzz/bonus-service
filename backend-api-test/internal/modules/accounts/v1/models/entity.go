package models

import (
	"github.com/google/uuid"
	"time"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
)

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UUID      uuid.UUID  `json:"uuid" gorm:"type:uuid;not null; unique"`
	Email     string     `json:"email" gorm:"not null; unique;"`
	Hash      string     `json:"hash" gorm:"varchar(255);default: null;"`
	Salt      string     `json:"hash" gorm:"varchar(255);default: null;"`
	FullName  string     `json:"name" gorm:"varchar(250);not null;"`
	Gender    GenderType `json:"gender" gorm:"type:gender_type; not null"`
	Photo     string     `json:"photo"`
	LastLogin *time.Time `json:"last_login" gorm:"null"`
	CreatedAt *time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:false; not null;"`
}
