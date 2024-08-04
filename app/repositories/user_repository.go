package repositories

import (
	"github.com/go-devops-demo/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindById(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Delete(user models.User) error {
	err := r.db.Delete(&user).Error
	return err
}
