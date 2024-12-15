package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// cleaning up stuff with 'defer'. Normally, a function call runs immediately, but defer
	// delays the invocation until the surrounding function exists.
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
