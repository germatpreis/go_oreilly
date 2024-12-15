# Links

* codebase for the book: https://github.com/learning-go-book-2e
* similar to npmjs: https://pkg.go.dev/

# Commands

```
go get ./.. - scans my source code to find out which modules to download 

go get 

go get github.com/shopspring/decimal

go get github.com/learning-go-book-2e/simpletax@v1.0.0 - get a specific version 
of a module

go get -u=patch - upgrade to the bug patch release

go get -u github.com/learning-go-book-2e/simpletax - upgrade to the latest 
available version

go mod tidy - scan the source code and synchronize the go.mod and go.sum files, 
adding and removing references

go list -m -versions github.com/learning-go-book-2e/simpletax - lists available 
version for a given module

go mod vendor - creates a directory 'vendor' at the top level of the module 
that contains all your modules dependencies. Needs to be run everythime a new 
dependency is added or upgraded. This has fallen out of favour. Can speed up builds though.
```

# Book

## Chapter 6: Pointers

How to create pointers:

```go
package main

var x int32 = 10
var y bool = true

pointerX := &x
pointerY := &y

var pointerZ *string
```

> The & is the address operator. It precedes the value type and returns the address
> where the value is stored:

```go
x := "hello"
pointerToX := &x
```

> the * is the indirection operator. It precedes a variable of pointer type and
> returns the pointed-to value. This is called dereferencing.

```go
package main

x := 10
pointerToX := &x
fmt.Println(pointerToX) // prints memory address
fmt.Println(*pointerToX) // prints 10
z := 5 + *pointerToX
fmt.Println(z) // prints 15
```

> Before dereferencing a pointer, you must make sure that the pointer is non-nil. Your
> program will panic if you attempt to dereference a `nil` pointer!

```go
package main

var x *int
fmt.Println(x == nil) // prints true
fmt.Println(*x) // panics
```

> A pointer type is a type that represents a pointer. It is written with a * before
> a type name. A pointer type can be based on any type.

```go
x := 10
var pointerToX *int
pointerToX := &x
```

> To create a pointer for structs do this:

```go
type person struct {
    firstName string
    lastName  string
}

mike := &person{
    firstName: "Mi",
    lastName:  "ke",
}

fmt.Println(mike) // prints &{Mi ke}
fmt.Println(&mike) // prints something like 0x14000116018
```

> when you need a pointer to a primitive type, declare a variable
> and point to it

```go
var y string
z := &y
```

> not being able to take the address of a constant is sometimes
> inconvenient. if you have a struct with a field of a pointer
> to a primitive type, you can't assign a literal directly to
> the field

```go
type person struct {
	FirstName string
	MiddleName *string
}

p1 := person{
	FirstName: "X",
	MiddleName: "Y" // this line won't compile
}

// solution
func makePointer[T any](t T) *T {
	return &t
}

p2 := person{
    FirstName: "X",
    MiddleName: makePointer("Y") // this works, see pg123 for explanation
}
```

From pg 124:

> **The difference between Go and other languages like javascript,
> java, etc is that Go gives you the > choice < to use pointers
> or values for both, primitives and structs. Most of the time you 
> should use a value. Values make it easier to understand how and
> when your data is modified. A secondary benefit is that using
> values reduces the amount of work that the garbage collector has
> to do**

From pg 125:

> The lack of immutable declarations in Go might seem problematic,
> but the ability to choose between value and pointer paraneter
> types addresses the issue. Using mutable objects is just fine
> if you are using them entirely locally within a method, and with
> only one reference to the object.
> 
> **Rather than declare that some variables and parameters are 
> immutable, Go Developers use pointers to indicate that a parameter is
> mutable.**
>
> **Since Go is a call-by-value language, the values passed to functions are 
> copies. For non-pointer types like primitives, structs, and arrays, 
> this means that the called function cannot modify the original. Since
> the called function cannot modify the original. SINCE THE CALLED FUNCTION
> HAS A COPY OF THE ORIGINAL DATA, THE ORIGINALS DATA IMMUTABILITY IS
> GUARANTEED**.

**NOTE: there are a few implications regarding the latest paragraph. Re-read 
pg125 and update the location here. IMPORTANT TO UNDERSTAND THIS!**

### Pointers are a last resort

> You should be careful when using pointers in Go. They make it harder
> to understand data flow and can create extra work for the garbage 
> collector.

> The only time you should use pointer parameters to modify a variable is
> when the function expects and interface ie when working with JSON support
> in Go's standard library.

### Pointer Passing Performance

See pg 130, but be aware that performance improvements are small - unless you are dealing with bigger data sizes.

### The Difference Between Maps and Slices

See pg 131, **very interesting theory regarding how slices and maps are working internally to go**. 

> You should be careful to consider maps for input parameters. Use structs whenever possible. Use maps only
> if the names of the keys are not known.

> By the default, you should assume that a slice is not modified by a function.

### Slices as buffers

See pg 135.

> Writing idiomatic Go means avoiding unneeded allocations. Rather than returning a new allocation each time your read
> from a data source, you create a slice of bytes once and use it as a buffer to read data from a source.

An example is shown how to read an process a file using a buffer of 100 bytes.

### Reducing the Garbage Collector's Workload

See pg 136, **very, very interesting theory regarding how stacks, frames, pointers and memory work together**.

## Chapter 7: Types, Methods and Interfaces

Example of custom types with methods.

```go
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
func (p Person) ToString() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Other examples of custom types

type Score int
type Converter func(string) Score
type TeamScores map[string]Score
```

### Pointer Receivers and Value Receivers

How to choose:

* if the method modifies the receiver, I **must** use a `pointer receiver`
* if the method needs to handle `nil` instances, then I **must** use a `pointer receiver`
* if the method doesn't modify the `receiver`, I **can** use a `value receiver`

It is idiomatic to only have `pointer receivers` on a type, as soon as there is a 
single one needed.

### Code your Methods for nil Instances

See pg 149 for a valid reason why where `nil pointer receivers` make sense (`IntTree`), 
most of the time it isn't though.

### iota is for Enumerations - Sometimes

See pg 152. This is to esoteric - skipped!

### Use Embedding for Composition

There is built in support for composition and promotion:

```go
package main

import "fmt"

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
```

Note that
```
fmt.Println(m.Description()) // prints Bob Bobson (1234)
```

can reference the field `ID` directly from the variable.

### A quick Lesson on interfaces

See pg 157. This is an interface

```go
package main

type Incrementer interface {
	Increment()
}
```

### Interfaces are Type-Safe Duck Typing

Here an example how interfaces are used together with structs:

```go
package main

type LogicProvider struct{}

func (l LogicProvider) Process(_ string) string {
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.Process("...data..")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
```

Aaaaaand interfaces can also be embedded:

```go
package main

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type ReadCloser interface {
	Reader
	Closer
}
```

### Accept Interfaces, Return Structs

See pg 162. Do as they say.

### Interfaces and nil

WTF??

### Interfaces are Comparable

WTF??

### The Empty Interface Says Nothing

There is something the need to say that a variable could store a value of `any` 
type, ie when reading json

```go
package main

data := map[string]any{}
contents, err := os.ReadFile("testdata/sample.json")
```

The recommendation though is to avoid `any` if possible.

### Type Assertions and Type Switches

TODO

## Chapter 10: Modules, Packages, and Imports

### Repositories, Modules and Packages

A `repository` is a place in a VCS. A `module` is a bundle of 
go source code that distributed and versioned as a single unit.
`Modules` are stored in a `repository`. `Modules` consist of one
or more `packages`, which are directories of source code. `packages`
give a `module` organization structure.

> While you can store more than one `module` in a `repository`, it
> is discouraged. Everything with a `module` is versioned together.
> Maintaining two modules in one repository requires you to track
> separate versions for two different modules in a single repository.

### Building Packages

> Go uses `captialization` to determine whether a pckage-level identifier
> is visible outside the package where it is declared. An identifier
> that starts with an **uppercase** is exported.
> 
> Anything you export is part of your package's API, Before you
> export an identifier, be sure that you intend to expose it to clients.
> Document  all exported identifiers and keep them backward compatible
> unless you are intentionally making a major version change.

### Naming conventions

pg 236

### How to structure your apps

https://www.youtube.com/watch?v=oL6JBUk6tj0
