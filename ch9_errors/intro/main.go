package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {
	remainder, mod, err := calcRemainderAndMod(10, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d %d\n", remainder, mod)

	produceSentinelError()
}

func produceSentinelError() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if errors.Is(err, zip.ErrFormat) {
		fmt.Println("Told you so")
	}
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}
