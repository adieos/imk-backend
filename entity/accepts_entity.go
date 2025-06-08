package entity

import (
	"github.com/google/uuid"
)

type Accept struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CategoryID   uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`
	BankSampahID uuid.UUID `gorm:"type:uuid;not null" json:"bank_sampah_id"`

	Quota  int    `json:"quota"`          // maximum amount of waste that can be accepted
	Filled int    `json:"filled"`         // current amount of waste accepted
	Note   string `json:"note,omitempty"` // optional note about the acceptance
	Price  int    `json:"price"`          // price per unit of waste accepted

	Category   *Category   `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	BankSampah *BankSampah `gorm:"foreignKey:BankSampahID" json:"bank_sampah,omitempty"`
}
