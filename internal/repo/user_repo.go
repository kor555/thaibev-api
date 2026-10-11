package repo

import (
	"api/internal/entity/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&domain.User{})
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserInfo(userID string) (string, error) {
	return "nil", nil
}

func (r *UserRepository) Save(users *domain.User) (*domain.User, error) {
	if err := r.db.Create(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
