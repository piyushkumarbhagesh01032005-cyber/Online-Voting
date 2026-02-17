package models

import "time"

type Vote struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	UserID      int       `gorm:"not null;index:idx_user_election,unique" json:"user_id" validate:"required"`
	ElectionID  int       `gorm:"not null;index:idx_user_election,unique" json:"election_id" validate:"required"`
	CandidateID int       `gorm:"not null" json:"candidate_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}
