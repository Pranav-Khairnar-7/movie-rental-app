package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:VARCHAR(36); PRIMARY_KEY;"`
	CreatedAt time.Time  `json:"createdAt" gorm:"type:DATETIME"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"type:DATETIME"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"type:DATETIME"`
}
