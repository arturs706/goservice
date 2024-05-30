package repository

import "example.com/gouserservice/domain"

type DBHandler interface {
	GetAllUsers() ([]*domain.User, error)
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(userID string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(userID string) error
	
}