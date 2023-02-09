package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

type AccountRepository interface {
	CreateAccount(username, password, email string) (int, error)
	ReadAccount(id int) (string, string, string, error)
	UpdateAccount(id int, username, password, email string) error
	DeleteAccount(id int) error
}

type AccountRepositoryImpl struct {
	accountRepository AccountRepository
}

func (r *AccountRepositoryImpl) CreateAccount(username, password, email string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO auth.accounts (username, password, email, created_on) VALUES ($1, $2, $3, NOW()) RETURNING user_id", username, password, email).Scan(&id)
	return id, err
}

func (r *AccountRepositoryImpl) ReadAccount(id string) (string, string, error) {
	var username, email string
	err := db.QueryRow("SELECT username, email FROM auth.accounts WHERE user_id = $1", id).Scan(&username, &email)
	return username, email, err
}

func (r *AccountRepositoryImpl) UpdateAccount(id string, username, password, email string) error {
	_, err := db.Exec("UPDATE auth.accounts SET username = $1, password = $2, email = $3 WHERE user_id = $4", username, password, email, id)
	return err
}

func (r *AccountRepositoryImpl) DeleteAccount(id string) error {
	_, err := db.Exec("DELETE FROM auth.accounts WHERE user_id = $1", id)
	return err
}
