package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// exercise two
	result, err := oneSimpleCalculator(10, 5, "/")
	if err != nil {
		log.Fatal("something weird happened: ", err)
	}
	fmt.Println("result is: ", result)

	// exercise three
	fn := prefixer(">>> ")
	prefixedResult := fn("cornholio")
	fmt.Println(prefixedResult) // >>> cornholio
}

func prefixer(prefixToBeUsed string) func(input string) string {
	return func(input string) string {
		return fmt.Sprintf("%s%s", prefixToBeUsed, input)
	}
}

func oneSimpleCalculator(a, b int, operand string) (int, error) {
	if operand != "/" {
		return 0, errors.New("only supported operand is '/'")
	}
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}
