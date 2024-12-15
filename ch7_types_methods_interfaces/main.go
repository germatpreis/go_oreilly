package main

import "fmt"

// Person should be read as 'user-defined type Person that has an
// UNDERLYING TYPE of struct literal that follows'
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// ToString methods can only be defined on package block level
// this > could < also be a value receiver
func (p *Person) ToString() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // this is an embedded field - no name!
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// logic
	return []Employee{}
}

func main() {
	p := Person{
		FirstName: "x",
		LastName:  "x",
		Age:       0,
	}
	fmt.Println(p.ToString())

	p.UpdateAge(24)
	fmt.Println(p.ToString())

	// code your methods for nil instances
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false

	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "1234",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)            // prints 1234
	fmt.Println(m.Description()) // prints Bob Bobson (1234)
}
