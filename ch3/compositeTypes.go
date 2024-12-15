package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	slicesBasics()
	slicesNil()
	slicesEqual()
	slicesLength()
	slicesAppend()
	slicesCreateWithInitialLengthAndCapacity()
	slicesClear()
	slicesSlicing()
	slicesCopying()
	slicesCopyingSubset()

	stringBasics()
	stringConvertRuneOrByteToString()

	mapBasics()
	mapReadingAndWriting()
	mapCommaIdiom()
	mapDeleteClearCompareEntries()

	setFakeItWithAMap()

	structBasics()
	structAnonymous()
	structsCompare()
}

func structsCompare() {
	type computer struct {
		numberCores int
	}

	pc1 := computer{3}
	pc2 := computer{3}
	pc3 := computer{4}
	fmt.Println(pc1 == pc2) // true
	fmt.Println(pc1 == pc3) // false
}

func structAnonymous() {
	// anonymous structs useful for marshaling and unmarshalling of data, or
	// for use in tests.
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet) // {Fido dog}
}

func structBasics() {
	type person struct {
		name string
		age  int
		pet  string
	}

	// first struct literal format to create bob
	// all properties need to be specified
	bob := person{
		"bob",
		62,
		"fluffy",
	}

	// second struct literal format support only
	// populating certain fields, in any order
	bob1 := person{
		name: "bob",
	}

	// print the struct or just a property
	fmt.Println(bob)       // {bob 62 fluffy}
	fmt.Println(bob1.name) // bob

}

func setFakeItWithAMap() {
	set := map[int]bool{}
	vals := []int{5, 10, 205}
	for _, v := range vals {
		set[v] = true
	}
	fmt.Println(len(vals), len(set)) // 3 3
	fmt.Println(set[5])              // true
	fmt.Println(set[11])             // false
}

func mapDeleteClearCompareEntries() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	isEqual := maps.Equal(m1, m2)
	fmt.Println(isEqual) // true

	delete(m1, "a")
	fmt.Println(m1) // map[b:2]

	clear(m1)
	fmt.Println(m1) // map[]

}

func mapCommaIdiom() {
	// aka test if a key exists
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true

	v, ok = m["nada"]
	fmt.Println(v, ok) // 0 false
	fmt.Println(v, ok) // 0 false
}

func mapReadingAndWriting() {
	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"]) // 1

	totalWins["Lions"]++
	fmt.Println(totalWins["Lions"]) // 3
}

func mapBasics() {
	// map[keyType]valueType
	// this is an empty map literal, this can be written to
	totalWins := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	someCounters := map[string]bool{
		"yes": true,
		"no":  false,
	}

	// more complex example, here the value is a slice of strings
	teams := map[string][]string{
		"Orcas": []string{"Fred", "Ralph"},
		"Lions": []string{"Sarah", "Peter"},
	}
	fmt.Println(totalWins, someCounters, teams)
}

func stringConvertRuneOrByteToString() {
	var a rune = 'x'
	var s = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s, s2) // x y
}

func stringBasics() {
	var s string = "hello there"
	var b = s[6]   // this is a byte
	fmt.Println(b) // 116
}

func slicesCopyingSubset() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	fmt.Println(y, num) // [1 2] 2
}

func slicesCopying() {
	// if I need a slice thats independnt of the original
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num) // [1 2 3 4] 4
}

func slicesSlicing() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	fmt.Println(y) // [1 2]
}

func slicesClear() {
	x := []int{3}
	fmt.Println(x) // [3]
	clear(x)
	fmt.Println(x) // [0] <!-- values only set to 0, length is kept!
	// NOTE: when slicing a slice, no copy is made! there are now 2 variables that
	// NOTE: share the same memory space. changes to one slice, change also the other.
	// NOTE: to avoid this, make sure that append doesn't cause an overwrite by using
	// NOTE: the FULL SPLICE EXPRESSION WHICH USES A THIRD PART, WHICH INDICATES THE LAST
	// NOTE: POSITION IN THE PARENTS SLICE CAPACITY THAT AVAILABLE FOR THE SUBSLICE

	// NOTE BEST PRACTICE: BE CAREFUL WHEN TAKING A SLICE FROM A SLICE. AVOID MODIFYING
	// NOTE BEST PRACTICE: SLICES AFTER THEY HAVE BEEN SLICED OR IF THEY WERE PRODUCED BY SLICING.
	// NOTE BEST PRACTICE: USE A THREE-PART SLICE EXPRESSION TO PREVENT APPEND FROM SHARING CAPACITY
	// NOTE BEST PRACTICE: BETWEEN SLICES.
}

func slicesCreateWithInitialLengthAndCapacity() {
	// creates a slice with the length of 0 and capacity of 10
	x := make([]int, 0, 10)
	x = append(x, 0, 1)
	fmt.Println(x) // [0 1]
}

func slicesAppend() {
	// append multiple values to an existing slice
	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x) // [1 2 3 4 5 6]

	// concat to slices
	y := []int{7}
	y = append(y, x...)
	fmt.Println(y) // [7 1 2 3 4 5 6]
}

func slicesLength() {
	a := []string{"a", "b", "c"}
	length := len(a)
	fmt.Println(length) // 3
}

func slicesEqual() {
	x := []int{1, 2, 3}
	z := []int{1, 2, 3}
	same := slices.Equal(x, z)
	fmt.Println(same) // true
}

func slicesNil() {
	var x []int
	var isNil = x == nil
	fmt.Println(isNil) // true
}

func slicesBasics() {
	// makes a slice
	var x = []int{10, 20, 30}
	fmt.Println(x) // [10 20 30]
	// makes an array (don't do this)
	var y = [...]int{10, 20, 30}
	fmt.Println(y) // [10 20 30]
	// with specifing indices
	var z = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z) // [1 0 0 0 0 4 6 0 0 0 100 15]

	z[0] = 100
	fmt.Println(z) // [100 0 0 0 0 4 6 0 0 0 100 15]
}
