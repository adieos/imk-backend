package entity

import (
	"github.com/adieos/imk-backend/helpers"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	ProvinceID  uuid.UUID `json:"province_id"`
	Instansi    string    `json:"instansi"`
	NoTelp      string    `json:"no_telp"`
	InfoFrom    string    `json:"info_from"`
	Jenjang     string    `json:"jenjang"`
	Role        string    `json:"role"`
	IsVerified  bool      `json:"is_verified"`
	ILJMajor    string    `json:"ilj_major"`
	ILJSubmajor string    `json:"ilj_submajor"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	// u.ID = uuid.New()
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
