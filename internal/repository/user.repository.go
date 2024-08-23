package repository

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=mock/mock_user_repository.go -source=user.repository.go
type UserRepository interface {
	Create(db *gorm.DB, user *entity.User) error
	CountByUsername(db *gorm.DB, username string) (int64, error)
	FindById(db *gorm.DB, user *entity.User, id int64) error
}

type userRepository struct {
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *userRepository {
	return &userRepository{
		Log: log,
	}
}

func (r *userRepository) CountByUsername(db *gorm.DB, username string) (int64, error) {
	var total int64
	var user entity.User
	err := db.Model(&user).Where("username = ?", username).Count(&total).Error
	return total, err
}

func (r *userRepository) Create(db *gorm.DB, user *entity.User) error {
	return db.Create(user).Error
}

func (r *userRepository) FindById(db *gorm.DB, user *entity.User, id int64) error {
	return db.Where("id = ?", id).First(user).Error
}
