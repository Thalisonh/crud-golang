package user

import "github.com/Thalisonh/crud-golang/database/entity"

type IUserService interface {
	GetUsers() (*[]entity.User, error)
	GetUser(userId int64) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) error
	UpdateUser(userId int64, user *entity.User) (*entity.User, error)
	//AddBook(userId int64, book *entity.Book) (*entity.User, error)
}

type ServiceUser struct {
	repository IUserRepository
}

func NewUserService(repository IUserRepository) IUserService {
	return &ServiceUser{repository: repository}
}

func (s *ServiceUser) GetUsers() (*[]entity.User, error){
	var users *[]entity.User
	users, err := s.repository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
func (s *ServiceUser) GetUser(userId int64) (*entity.User, error){
	var user *entity.User

	user, err := s.repository.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *ServiceUser) CreateUser(user *entity.User) (*entity.User, error) {
	user, err := s.repository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, err
}
func (s *ServiceUser) DeleteUser(user *entity.User) error {
	//user, err := s.repository.GetUser(int64(user.ID))
	//if err != nil {
	//	return err
	//}

	bookError := s.repository.DeleteUser(user)
	if bookError != nil {
		return bookError
	}

	return nil
}

func (s *ServiceUser) UpdateUser(userId int64, user *entity.User) (*entity.User, error){
	user, err := s.repository.UpdateUser(userId, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
