package repository

import (
	"context"
	"errors"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/entity"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error)
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		DeleteUser(ctx context.Context, tx *gorm.DB, userId string) error
		ResetPassword(context.Context, string, string) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, tx *gorm.DB, userId string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.User{}, "id = ?", userId).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ResetPassword(ctx context.Context, email, hashedPassword string) error {
	// query email apakah ada
	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ErrEmailNotFound
		}
		return err
	}

	// ketemu maka update
	if err := r.db.WithContext(ctx).Model(&entity.User{}).Where("email = ?", email).Update("password", hashedPassword).Error; err != nil {
		return err
	}
	return nil
}
