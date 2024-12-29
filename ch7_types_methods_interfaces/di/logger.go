package main

import "fmt"

type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

type Logger interface {
	Log(message string)
}

func LogOutput(message string) {
	fmt.Println(message)
}
