package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
	"example.com/gouserservice/domain"
	"fmt"
)

type DBHandler struct {
	Conn *sql.DB
}

func NewDBHandler() DBHandler {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	return DBHandler{Conn: db}
}
func (dbHandler DBHandler) GetAllUsers() ([]*domain.User, error) {
	rows, err := dbHandler.Conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		user := new(domain.User)
		err := rows.Scan(&user.UserID, &user.FullName, &user.DOB, &user.Gender, &user.MobPhone, &user.Email, &user.EmailVerified, &user.EmailVerToken, &user.Passwd, &user.AuthMethod, &user.SocialID, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (dbHandler DBHandler) CreateUserDB(user *domain.User) error {
    stmt, err := dbHandler.Conn.Prepare("INSERT INTO users(user_id, full_name, dob, gender, mob_phone, email, email_ver, email_ver_token, passwd, auth_method, social_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *")
    if err != nil {
        return err
    }
    defer stmt.Close()

    err = stmt.QueryRow(user.UserID, user.FullName, user.DOB, user.Gender, user.MobPhone, user.Email, user.EmailVerified, user.EmailVerToken, user.Passwd, user.AuthMethod, user.SocialID, user.CreatedAt).
        Scan(&user.UserID, &user.FullName, &user.DOB, &user.Gender, &user.MobPhone, &user.Email, &user.EmailVerified, &user.EmailVerToken, &user.Passwd, &user.AuthMethod, &user.SocialID, &user.CreatedAt)
    if err != nil {
        return err
    }

    return nil
}

func (dbHandler DBHandler) GetUserByEmail(email string) (*domain.User, error) {
	row := dbHandler.Conn.QueryRow("SELECT * FROM users WHERE email=$1", email)
	user := new(domain.User)
	err := row.Scan(&user.UserID, &user.FullName, &user.DOB, &user.Gender, &user.MobPhone, &user.Email, &user.EmailVerified, &user.EmailVerToken, &user.Passwd, &user.AuthMethod, &user.SocialID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}




func (dbHandler DBHandler) GetUserByID(userID string) (*domain.User, error) {
	row := dbHandler.Conn.QueryRow("SELECT * FROM users WHERE user_id=$1", userID)
	user := new(domain.User)
	err := row.Scan(&user.UserID, &user.FullName, &user.DOB, &user.Gender, &user.MobPhone, &user.Email, &user.EmailVerified, &user.EmailVerToken, &user.Passwd, &user.AuthMethod, &user.SocialID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (dbHandler DBHandler) UpdateUser(user *domain.User) error {
	stmt, err := dbHandler.Conn.Prepare("UPDATE users SET full_name=$1, dob=$2, gender=$3, mob_phone=$4, email=$5, email_ver=$6, email_ver_token=$7, passwd=$8, auth_method=$9, social_id=$10, updated_at=$11 WHERE user_id=$12")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FullName, user.DOB, user.Gender, user.MobPhone, user.Email, user.EmailVerified, user.EmailVerToken, user.Passwd, user.AuthMethod, user.SocialID, time.Now(), user.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) DeleteUser(userID string) error {
	stmt, err := dbHandler.Conn.Prepare("DELETE FROM users WHERE user_id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) Close() error {
	err := dbHandler.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}


func (dbHandler DBHandler) LoginUserDB(email string, passwd string) (*domain.User, error) {
	row := dbHandler.Conn.QueryRow("SELECT * FROM users WHERE email=$1", email)
	user := new(domain.User)
	err := row.Scan(&user.UserID, &user.FullName, &user.DOB, &user.Gender, &user.MobPhone, &user.Email, &user.EmailVerified, &user.EmailVerToken, &user.Passwd, &user.AuthMethod, &user.SocialID, &user.CreatedAt)
	if err != nil {
		return nil, err
	} else {

		return user, nil

	}
	
}

