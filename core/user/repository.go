package user

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUser (userId int64) (*entity.User, error)
	GetUsers() (*[]entity.User, error)
	UpdateUser (userId int64, user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) error
	CreateUser (user *entity.User) (*entity.User, error)
}

type RepositoryUser struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &RepositoryUser{db: db}
}

func (r *RepositoryUser) GetUser (userId int64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", userId).First(&user).Error

	return &user, err
}

func (r *RepositoryUser) GetUsers() (*[]entity.User, error) {
	var user []entity.User
	err := r.db.Find(&user).Error

	return &user, err
}

func (r *RepositoryUser) UpdateUser (userId int64, user *entity.User) (*entity.User, error) {
	return user, r.db.Where("id = ?", userId).Save(&user).Error
}

func (r *RepositoryUser) DeleteUser(user *entity.User) error {
	return r.db.Delete(&user).Error
}

func (r *RepositoryUser) CreateUser (user *entity.User) (*entity.User, error) {
	return user, r.db.Create(&user).Error
}


