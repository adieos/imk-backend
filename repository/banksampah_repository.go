package repository

import (
	"context"
	"errors"

	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/entity"
	"gorm.io/gorm"
)

type (
	BSRepository interface {
		CreateBS(ctx context.Context, tx *gorm.DB, bs entity.BankSampah) (entity.BankSampah, error)
		GetBSById(ctx context.Context, tx *gorm.DB, bsId string) (entity.BankSampah, error)
		GetAllBS(ctx context.Context, tx *gorm.DB) ([]entity.BankSampah, error)
		GetAllBSByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.BankSampah, error)
		UpdateBS(ctx context.Context, tx *gorm.DB, bs entity.BankSampah) (entity.BankSampah, error)
		ChangeStatusBS(ctx context.Context, tx *gorm.DB, bsId string, status string) (entity.BankSampah, error)
		GetBSAccepts(ctx context.Context, tx *gorm.DB, bsId string) ([]entity.Accept, error)
		CreateBSAccept(ctx context.Context, tx *gorm.DB, accept entity.Accept) (entity.Accept, error)
	}

	bsRepository struct {
		db *gorm.DB
	}
)

func NewBSRepository(db *gorm.DB) BSRepository {
	return &bsRepository{
		db: db,
	}
}

func (r *bsRepository) CreateBS(ctx context.Context, tx *gorm.DB, bs entity.BankSampah) (entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&bs).Error; err != nil {
		return entity.BankSampah{}, err
	}

	return bs, nil
}

func (r *bsRepository) GetBSById(ctx context.Context, tx *gorm.DB, bsId string) (entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	var bs entity.BankSampah
	if err := tx.WithContext(ctx).Where("id = ?", bsId).Take(&bs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.BankSampah{}, dto.ErrBankSampahNotFound
		}
		return entity.BankSampah{}, err
	}

	return bs, nil
}

func (r *bsRepository) GetAllBS(ctx context.Context, tx *gorm.DB) ([]entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	var bss []entity.BankSampah
	if err := tx.WithContext(ctx).Find(&bss).Error; err != nil {
		return nil, err
	}

	return bss, nil
}

func (r *bsRepository) GetAllBSByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	var bss []entity.BankSampah
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).Find(&bss).Error; err != nil {
		return nil, err
	}

	if len(bss) == 0 {
		return nil, dto.ErrBankSampahNotFound
	}

	return bss, nil
}

func (r *bsRepository) UpdateBS(ctx context.Context, tx *gorm.DB, bs entity.BankSampah) (entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&bs).Error; err != nil {
		return entity.BankSampah{}, err
	}

	return bs, nil
}

func (r *bsRepository) ChangeStatusBS(ctx context.Context, tx *gorm.DB, bsId string, status string) (entity.BankSampah, error) {
	if tx == nil {
		tx = r.db
	}

	var bs entity.BankSampah
	if err := tx.WithContext(ctx).Model(&bs).Where("id = ?", bsId).Update("status", status).Error; err != nil {
		return entity.BankSampah{}, err
	}

	if err := tx.WithContext(ctx).Where("id = ?", bsId).Take(&bs).Error; err != nil {
		return entity.BankSampah{}, err
	}

	return bs, nil
}

func (r *bsRepository) GetBSAccepts(ctx context.Context, tx *gorm.DB, bsId string) ([]entity.Accept, error) {
	if tx == nil {
		tx = r.db
	}

	var accepts []entity.Accept
	if err := tx.WithContext(ctx).Where("bank_sampah_id = ?", bsId).Find(&accepts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dto.ErrBankSampahNotFound
		}
		return nil, err
	}

	return accepts, nil
}

func (r *bsRepository) CreateBSAccept(ctx context.Context, tx *gorm.DB, accept entity.Accept) (entity.Accept, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&accept).Error; err != nil {
		return entity.Accept{}, err
	}

	return accept, nil
}
