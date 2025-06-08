package entity

import (
	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name    string `json:"name"`
	LogoURL string `json:"logo_url,omitempty"` // optional URL for the category logo
}
