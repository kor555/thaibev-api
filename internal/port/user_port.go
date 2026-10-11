package port

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
)

type UserService interface {
	GetUserInfo(userID string) (string, error)
	SaveUserInfo(user model.User) (*domain.User, error)
}

type UserRepository interface {
	GetUserInfo(userID string) (string, error)
	Save(users *domain.User) (*domain.User, error)
}
