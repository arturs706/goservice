package domain

import (
    "context" 
    "time" 
)


type User struct {
    UserID        string    `json:"user_id,omitempty"`
    FullName      string    `json:"full_name,omitempty"`
    DOB           string    `json:"dob,omitempty"`
    Gender        string    `json:"gender,omitempty"`
    MobPhone      string    `json:"mob_phone,omitempty"`
    Email         string    `json:"email"`
    EmailVerified bool      `json:"email_ver,omitempty"`
    EmailVerToken string    `json:"email_ver_token,omitempty"`
    Passwd        string    `json:"passwd,omitempty"`
    AuthMethod    string    `json:"auth_method,omitempty"`
    SocialID      string    `json:"social_id,omitempty"`
    CreatedAt     time.Time `json:"created_at,omitempty"`
}

type UserRepository interface {
    GetAllUsers(ctx context.Context) ([]*User, error)
    CreateLocal(ctx context.Context, user *User) error
    GetByEmail(ctx context.Context, email string) (*User, error)
    GetByID(ctx context.Context, userID string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, userID string) error

}