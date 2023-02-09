package service

import (
	"fmt"
	"postgresql/repository"
)

var s = repository.AccountRepositoryImpl{}

func AccountCreate(username string, password string, email string) int {
	createdID, err := s.CreateAccount(username, password, email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Account created with ID: %d\n", createdID)
	return createdID
}

func AccountRead(userID string) (string, string) {
	username, email, err := s.ReadAccount(userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Account details: %s, %s, %s\n", username, email)
	return username, email
}

func AccountUpdate(accountID string, username string, password string, email string) {
	err := s.UpdateAccount(accountID, username, password, email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account updated.")
}

func AccountDelete(createdID string) {
	err := s.DeleteAccount(createdID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account deleted.")
}
