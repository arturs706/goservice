package repository

import (
    "context"
    "example.com/gouserservice/domain"
	"fmt"
)

type UserRepo struct {
    handler DBHandler
}

func NewUserRepo(handler DBHandler) UserRepo {
    return UserRepo{handler}
}

func (repo UserRepo) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
    users, err := repo.handler.GetAllUsers()
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (repo UserRepo) CreateLocal(ctx context.Context, user *domain.User) error {
    fmt.Println("Create user user-repository.go")
    err := repo.handler.CreateUserDB(user)
    if err != nil {
        return err
    }
    return nil
}

func (repo UserRepo) CreateGoogle(ctx context.Context, user *domain.User) error {
    fmt.Println("Create user user-repository.go")
    err := repo.handler.CreateUserDB(user)
    if err != nil {
        return err
    }
    return nil
}
func (repo UserRepo) CreateFacebook(ctx context.Context, user *domain.User) error {
    fmt.Println("Create user user-repository.go")
    err := repo.handler.CreateUserDB(user)
    if err != nil {
        return err
    }
    return nil
}


func (repo UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    user, err := repo.handler.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (repo UserRepo) GetByID(ctx context.Context, userID string) (*domain.User, error) {
    user, err := repo.handler.GetUserByID(userID)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (repo UserRepo) Update(ctx context.Context, user *domain.User) error {
    err := repo.handler.UpdateUser(user)
    if err != nil {
        return err
    }
    return nil
}

func (repo UserRepo) Delete(ctx context.Context, userID string) error {
    err := repo.handler.DeleteUser(userID)
    if err != nil {
        return err
    }
    return nil
}

func (repo UserRepo) LoginUserDomain(ctx context.Context, email string) (*domain.User, error) {
	fmt.Println("Login user user-repository.go")
    user, err := repo.handler.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }
    return user, nil
}