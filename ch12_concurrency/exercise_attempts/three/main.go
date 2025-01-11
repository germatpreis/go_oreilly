package main

import (
	"fmt"
	"math"
)
import "sync"

var cache = map[int]float64{}
var initCacheOnce = sync.OnceFunc(initCache)

func initCache() {
	fmt.Println("Init cache")
	for i := 0; i < 100_000; i++ {
		cache[i] = math.Sqrt(float64(i))
	}
}

func lookup(key int) float64 {
	initCacheOnce()
	return cache[key]
}

func main() {
	fmt.Println(lookup(1))
	fmt.Println(lookup(2000))
}
