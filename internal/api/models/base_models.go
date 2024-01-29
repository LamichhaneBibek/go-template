package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int       `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	UpdatedAt time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = time.Now().UTC()
	return
}
