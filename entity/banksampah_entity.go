package entity

import (
	"github.com/google/uuid"
)

type BankSampah struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	Name        string `json:"name"`
	Address     string `json:"address"`
	Contact     string `json:"contact"`
	Status      string `json:"status"`                // "active", "inactive", etc.
	AcceptAll   bool   `json:"accept_all"`            // true if accepts all types of waste, false otherwise
	OpenHours   string `json:"open_hours"`            // e.g., "Mon-Fri 08:00-17:00"
	Description string `json:"description,omitempty"` // optional description of the bank sampah
	PhotoURL    string `json:"photo_url,omitempty"`   // optional URL for a photo of the bank sampah
	City        string `json:"city,omitempty"`        // optional city where the bank sampah is located

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
