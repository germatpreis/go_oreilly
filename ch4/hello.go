package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ifWithScopedVariable()

	// when to pick which for loop construct
	// - by default use the for-range. it works well for strings (returns runes, not bytes), slices
	//   maps and it works well with channels
	// - use a standard loop if you don't want to start at 0

	forFlavourOne()
	forFlavourTwo()
	//forFlavourThree()
	forFlavourFour()

	switchIntro()

	exerciseOneAndTwo()
	exerciseThree()
}

func exerciseThree() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
}

func exerciseOneAndTwo() {
	var result []int
	for i := 0; i < 100; i++ {
		result = append(result, rand.Intn(100))
	}

	for _, val := range result {
		if val%2 == 0 {
			fmt.Println("two")
		} else if val%3 == 0 {
			fmt.Println("three")
		} else if val%2 == 0 && val%3 == 0 {
			fmt.Println("six")
		} else {
			fmt.Println("nevermind")
		}
	}
}

func switchIntro() {
	var words []string
	words = append(words, "a", "cow", "smile", "gopher")

	// switch with integer
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length", wordLen)
		case 6, 7, 8, 9:
			fmt.Println(word, "is a long word!")
		}
	}

	// switch with expression that evaluates to a boolean
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "is exactly the right length")
		}
	}
}

func forFlavourFour() {
	// for-range loop over a slice
	evenVals := []int{2, 4, 5, 6, 7}
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// for-range loop over a sett'ish map (not interested in the value, thus can be ommitted)
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

}

func forFlavourThree() {
	// this is an infinite loop
	for {
		fmt.Println("hello")
	}
}

func forFlavourTwo() {
	// this is basically a while
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func forFlavourOne() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func ifWithScopedVariable() {
	// scoping a variable to an if statement
	// its only available in the if / else if / else blocks
	if n := rand.Intn(10); n == 0 {
		fmt.Println("thats too low")
	} else if n > 5 {
		fmt.Println("thats too big")
	} else {
		fmt.Println("thats a good number")
	}
}
