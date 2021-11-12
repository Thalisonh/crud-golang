package user

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

var db *gorm.DB

type IUserRepository interface {
	GetUser (userId int64) (*entity.User, error)
	GetUsers() (*[]entity.User, error)
	UpdateUser (userId int64, user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) error
	CreateUser (user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUser (userId int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", userId).First(&user).Error

	return &user, err
}

func (r *UserRepository) GetUsers() (*[]entity.User, error) {
	var user []entity.User
	err := r.db.Find(&user).Error

	return &user, err
}

func (r *UserRepository) UpdateUser (userId int64, user *entity.User) (*entity.User, error) {
	return user, r.db.Where("id = ?", userId).Save(&user).Error
}

func (r *UserRepository) DeleteUser(user *entity.User) error {
	return r.db.Where("id = ?", user.ID).Delete(&entity.Book{}).Error
}

func (r *UserRepository) CreateUser (user *entity.User) (*entity.User, error) {
	return user, r.db.Create(&user).Error
}


