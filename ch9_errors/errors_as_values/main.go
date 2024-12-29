package main

import (
	"errors"
	"fmt"
	"os"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found"),
			Err:     err,
		}
	}
	return data, nil
}

func getData(token string, file string) ([]byte, error) {
	return nil, errors.New("couldn't get data")
}

func login(uid string, pwd string) (string, error) {
	return "", errors.New("foobar")
}

func main() {
	data, err := LoginAndGetData("xxx", "xxx", "xxx")
	if err != nil {
		fmt.Println("woopsie")
		os.Exit(1)
	}
	fmt.Println(data)
}
