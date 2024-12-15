package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	person       Person
	employeeId   string
	accessRights map[string]bool
}

func main() {
	//one()
	//two()
	testStructs()
}

func testStructs() {
	slartibartfast := Person{
		name: "slarti",
		age:  22,
	}
	emp := Employee{
		person:     slartibartfast,
		employeeId: "A234234",
		accessRights: map[string]bool{
			"fullProd": true,
		},
	}
	fmt.Println(emp)
}

func two() {
	const value = 10
	const i int = value
	const f float64 = value
	fmt.Println(i, f)
}

func one() {
	i := 20
	fmt.Println(i)
}
