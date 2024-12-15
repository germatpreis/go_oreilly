package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {
	// basic function that has two consecutive params of the same type, need only
	// specify the type in the last param
	div(12, 1)

	// pattern to simulate named / optional parameters
	// a function shouldn't have to many parameters and named & optional ones
	// are mostly used if you have a lot. this might be a code smell.
	MyFunc(MyFuncOpts{FirstName: "flobert", LastName: "cornholio", Age: 18})

	// varadic input paramters also exist
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))

	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	// multiple return values, ignored ones should be set to _
	result, _, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	// multiple return values, again, but the return values are named
	// this 'pre-declares' variables that can be used within the function
	// they have potentially some weird sideeffects (pg 100), so maybe
	// not use them
	result, _, err = divAndRemainderNamedReturnValues(5, 2)

	// functions can also be used as variables
	var myFuncVariable func(string) int
	myFuncVariable = f1
	result = myFuncVariable("Hello")
	fmt.Println(result)

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)

	// anonymous functions
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 10; i++ {
		f(i)
	}

	// passing functions as parameters (here: sorting function, the people slice
	// is sorted 'in place' (go is 'call by value')
	passingFunctionsAsParameter()

	// returning a function
	fn := returningFunctionFromAFunction(7)
	var fnCallResult = fn(12)
	fmt.Println(fnCallResult)

	// cleaning up stuff with 'defer'. Normally, a function call runs immediately, but defer
	// delays the invocation until the surrounding function exists.
	// MOVED TO ch5_defer_showcase/MyCat.go BECAUSE IT WOULD INTERFERE WITH THE OTHER PROGRAMS HERE

	// 'defer': LIFO (last in, first out)
	// the code within defer runs AFTER the return statement. the input parameters are evaluated
	// immediately and their values are stored until the function runs.
	//
	// there is a way for a deferred function to examine or modify the return values of its surrounding
	// function. there is, and its the best reason to use named return values. it allows your code to
	// take actions based on an error (ie. database transaction cleanup)
	//
	// a common pattern in Go is for the function that allocates a resource to also return a closure
	// that cleans up the resource
	deferFiFoExample()

	// go is CALL BY VALUE. it means that when you supply a variable for a parameter to a function, go
	// always MAKES A COPY OF THE VALUE OF THE VARIABLE. This is true for primitive types and structs. for
	// maps and slices this is a bit different (they are implemented via pointers)
	showcaseCallByValuePrimitivesAndStruct()
	showcasePointersMapsAndSlices()
}

func showcasePointersMapsAndSlices() {
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	s := []int{1, 2, 3}
	modSlice(s)

	fmt.Println(m, s)
}

func modSlice(s []int) {
	// this will change the slice
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func modMap(m map[int]string) {
	// this will change the map
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

type Person struct {
	age  int
	name string
}

func showcaseCallByValuePrimitivesAndStruct() {
	p := Person{22, "internal"}
	i := 2
	s := "hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

func modifyFails(i int, s string, p Person) {
	i = i * 2
	s = "Goodbye"
	p.name = "bob"
}

func deferFiFoExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first: ", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("second: ", val)
	}(a)

	a = 30
	fmt.Println("exiting: ", a)

	// exiting:  30
	// second:  20
	// first:  10

	return a
}

func returningFunctionFromAFunction(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func passingFunctionsAsParameter() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaugther", 23},
	}
	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func div(num, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func divAndRemainderNamedReturnValues(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	//
}
