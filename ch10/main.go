package main

import (
	"ch10/format"
	"ch10/internal"
	"ch10/math"
	//crand "crypto/rand"
	"fmt"
	//mrand "math/rand"
	"github.com/learning-go-book-2e/formatter"
	"github.com/shopspring/decimal"
)

// main
func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)

	internal.Doubler(12)

	//mrand.Intn(5)
	//crand.Int()
	formatter.Warning("woha")
}
