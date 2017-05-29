package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username string, password string) *User {
	return &User{username, password}
}

func (u *User) Create() error {
	if _, err := os.Stat(fmt.Sprintf("%v.user", u.Username)); err == nil {
		return errors.New("user exists")
	}
	return ioutil.WriteFile(fmt.Sprintf("%v.user", u.Username), []byte(u.Password), 0644)
}

func (u *User) Get(username *string) error {
	password, err := ioutil.ReadFile(fmt.Sprintf("%s.user", *username))
	if err != nil {
		return err
	}
	u.Username = *username
	u.Password = string(password)
	return nil
}

func (u *User) Auth() (bool, error) {
	password, err := ioutil.ReadFile(fmt.Sprintf("%v.user", u.Username))
	if err != nil {
		return false, err
	}
	if u.Password == string(password) {
		return true, nil
	}
	return false, nil
}
