package repositories

import (
	"context"
	"time"
	"user-service/internal/infrastructure/db/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(client *gorm.DB) UserRepository {
	return &userRepository{db: client}
}

func (ur *userRepository) GetByID(ctx context.Context, id int) (model *models.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = ur.db.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}

	return
}

func (ur *userRepository) GetByPhone(ctx context.Context, phone string) (model *models.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = ur.db.WithContext(ctx).Where("phone = ?", phone).First(&model).Error
	if err != nil {
		return nil, err
	}

	return
}
