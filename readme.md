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

`type assertions` are used to retrieve the dynamic value of an interface. If the assertion
is invalid, it causes a runtime panic unless I use the `comma, ok` idiom. Note: there is no type
conversion going (the value isn't changed to a new type). Type assertions 'reveal' the type of the value
stored in the interface.

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    // Type assertion without checking
    s := i.(string)
    fmt.Println(s)

    // Type assertion with `comma, ok` idiom
    s, ok := i.(string)
    if ok {
        fmt.Println(s)
    } else {
        fmt.Println("Type assertion failed")
    }
}
```

And this is how a `type switch` looks like:

```
TODO
```

Note: `type assertions` abd `type switches` should be used **sparingly**. One use case where it makes sense
to use it is to check if the type behind the interface also implements another interface.

### Function Types are a Bridge to Interfaces

`Go` allows methods on *any* user-defined type, including user-defined function types. This is actually very
useful. They allow functions to implement interfaces. The most common usage is for HTTP handlers.

### Implicit Interfaces Make Dependency Injection Easier

See complete manual `DI` example at `ch7_types_methods_interfaces/di`.

`DI` is easy to implement in `Go` without any framework.

## Chapter 9: Errors

### Use strings for simple errors

There are two ways to create errors:

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	errOne := errors.New("hello")
	errTwo := fmt.Errorf("%d isn't an even number", 5)
}
```

### Sentinel Errors

Signal that processing cannot continue because of a problem with the current state. They are
one of the few variables tha are declared at the package level. By convention they start with
`Err`. They should be treated as read-only.

Once you define a `sentinel error` it is part of your public API, so make sure you really need
it!

Here is how you can check for a sentinel error (the book mentions actually using `==` but 
`goland` complains about it and proposes the variant below).

```go
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func produceSentinelError() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
```

### Errors are values

Here is how you can define you own `error` that transports more
than a single string (ie. a `Status` and the original `error`).

```go
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
```

### Wrapping Errors

Often when an errors is passed back through your code, you often
want to add information to it. To wrap an error, use `fmt.Errof`
together with the special `verb` `%w`. 

Usually you don't use `Unwrap` like below, but rather `errors.Is`
and `errors.As` to find a specific wrapped error.

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err) // here %w wraps the original error
		//return fmt.Errorf("in fileChecker: %v", err) // here %v only takes the message from the original error
	}
	f.Close()
	return nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}
```

### Wrapping Multiple Errors

Sometimes you need to return multiple errors (ie. when you validate a struct
and every validation error of multiple fields should be its own 
error). Since the standard function signature is `error` and not `[]error`,
I need to merge those errors using `errors.Join`.

```go
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
```

### Is and As

TODO: pg 214

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

## Chapter 11: Go Tooling

### Using `go run` to try out small programs

`go run` builds and executes a program in one step. The binary is created in
a temporary directory (which is deleted after executing the program).

### Adding Third-Party tools with `go install`

While some people choose to distribute their Go programs as precompiled 
binaries, tools written in Go can also be built from source and installed
on your computer via the `go install` command.

By default, `go install` places the binaries into the `go/bin` directory within
your home directory. It is strongly recommended to add this location to your `PATH`
(it isn't done by default).

```sh
go install github.com/rakyll/hey@latest
```

Updating is done the same way (aka at `@latest` pulls the last version).

`go install` has become the method of choice for distributing third-party
developer tools.

### Improve Import Formatting with goimports

This is the improved version of `go fmt`. It needs to be installed:

```sh
go install golang.org/x/tools/cmd/goimports@latest

# -l: print files with incorrect formatting
# -w: modify the files in place
goimports -l -w .
```

## Chapter 12: Concurrency

> For our purposes, all you need to understand is that more concurrency does
> not mean more speed.
> ...
> Use concurrency when you want to combine data from multiple operations that
> can operate independently.
> ...
> Another important thing to note is that concurrency isn't worth using if the
> process that's running concurrently doesn't take a lot of time. Concurrency
> isn't free. 
> ...
> This is why concurrent operations are often used for I/O; reading or writing 
> to disk or network is thousands times slower than all but he most complicated
> in-memory processes. If you are not sure if concurrency will help, first write
> your code serially and then write a benchmark to compare performance with a 
> concurrent implementation.

### Goroutines

The `goroutine` is the core concept in Go's concurrency model. Think of a `goroutine` as
a lightweight thread, managed by the `go runtime`. When a `go` program starts, the `go runtime`
create a number of threads and launches a single `goroutine` to run your program. All the
goroutines created by your program, including the initial one, are assigned to the OS threads 
automatically by the `go runtime scheduler`.

There are a lot of advantages to using `goroutines` and it allows `go programs` to spawn tens
of thousands of simultaneous goroutines.

A `goroutine` is launched by placing the `'go'` keyword before a function invocation.
