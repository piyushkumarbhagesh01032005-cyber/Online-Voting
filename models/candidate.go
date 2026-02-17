package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Id         int    `gorm:"primarykey" json:"id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=3,max=100"`
	ElectionId int    `gorm:"not null;index" json:"election_id" validate:"required"`
}
