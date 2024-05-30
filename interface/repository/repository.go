package repository

import "example.com/gouserservice/domain"

type DBHandler interface {
	GetAllUsers() ([]*domain.User, error)
	CreateUserDB(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(userID string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(userID string) error
	LoginUserDB(email string, password string) (*domain.User, error)
	
}