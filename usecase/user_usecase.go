package usecase

import (
	"database/sql"
	"restfull-api/m/v2/domain"
)

type UserRepository interface {
	GetUserByID(id int) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) (sql.Result, error)
	DeleteUser(id int) (sql.Result, error)
	ListUsers() ([]*domain.User, error)
}

type UserUseCase struct {
	UserRepo UserRepository
}

func (u *UserUseCase) GetUserByID(id int) (*domain.User, error) {
	return u.UserRepo.GetUserByID(id)
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
	return u.UserRepo.CreateUser(user)
}

func (u *UserUseCase) UpdateUser(user *domain.User) (sql.Result, error) {
	return u.UserRepo.UpdateUser(user)
}

func (u *UserUseCase) DeleteUser(id int) (sql.Result, error) {
	return u.UserRepo.DeleteUser(id)
}

func (u *UserUseCase) ListUsers() ([]*domain.User, error) {
	return u.UserRepo.ListUsers()
}
