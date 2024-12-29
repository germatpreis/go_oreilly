package main

import (
	"errors"
	"fmt"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) Validate() error {
	var errs []error

	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}

	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func main() {
	p := Person{}
	err := p.Validate()
	if err != nil {
		fmt.Printf("An error has happend\n\n%v", err)
		os.Exit(1)
	}
}
