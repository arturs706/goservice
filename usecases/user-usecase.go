package usecases

import (
    "context"
    "example.com/gouserservice/domain"
    "log"
)

type UserInteractor struct {
    UserRepository domain.UserRepository
}

func NewUserInteractor(repo domain.UserRepository) UserInteractor { 
    return UserInteractor{repo}
}


func (interactor *UserInteractor) CreateLocalUser(user *domain.User) error {
	err := interactor.UserRepository.CreateLocal(context.Background(), user) 
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (interactor *UserInteractor) GetUserByEmail(email string) (*domain.User, error) {
	user, err := interactor.UserRepository.GetByEmail(context.Background(), email) 
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (interactor *UserInteractor) GetUserByID(userID string) (*domain.User, error) {
	user, err := interactor.UserRepository.GetByID(context.Background(), userID) // Add context.Background() here
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (interactor *UserInteractor) UpdateUser(ctx context.Context, user *domain.User) error {
	err := interactor.UserRepository.Update(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (interactor *UserInteractor) DeleteUser(userID string) error {
	err := interactor.UserRepository.Delete(context.Background(), userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (interactor *UserInteractor) GetAllUsers() ([]*domain.User, error) {
	users, err := interactor.UserRepository.GetAllUsers(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return users, nil
}
