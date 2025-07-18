package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bioskop struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Nama      string    `json:"nama"`
	Lokasi    string    `json:"lokasi"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Bioskop) TableName() string {
    return "bioskops"
}

func (b *Bioskop) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}