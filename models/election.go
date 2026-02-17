package models

import (
	"time"

	"gorm.io/gorm"
)

type Election struct {
	gorm.Model
	Title     string    `gorm:"type:varchar(100);not null" json:"title" validate:"required,min=3,max=100"`
	StartTime time.Time `gorm:"not null" json:"start_time" validate:"required"`
	EndTime   time.Time `gorm:"not null" json:"end_time" validate:"required,gtfield=StartTime"`
	Status    string    `gorm:"type:varchar(50);default:'upcoming'" json:"status" validate:"required,oneof=upcoming active completed"`
}
