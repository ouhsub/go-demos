package proxy

import (
	"log"
	"time"
)

type IUser interface {
	Login(string, string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (up *UserProxy) Login(username, password string) error {
	start := time.Now()
	if err := up.user.Login(username, password); err != nil {
		return err
	}
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}
