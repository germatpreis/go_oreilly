package main

import (
	"fmt"
	"strings"
)

func main() {
	contains := strings.Contains("internal", "foo")
	fmt.Println(contains)

	var slice = []string{"her", "mitn", "bert"}
	slice = append(slice, "snowbert")
	fmt.Println(slice)

	var mapsi = map[string][]int{}
	mapsi["flow"] = []int{1, 2, 3}
	fmt.Println(mapsi)
}
