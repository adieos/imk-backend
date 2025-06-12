package migrations

import (
	"github.com/adieos/imk-backend/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.BankSampah{},
		&entity.Accept{},
	); err != nil {
		return err
	}

	return nil
}
