package processors

import (
	"context"
	"errors"
	"math/rand"
)

var prefix = "hello"

type AProcessor struct {
	in     chan string
	out    chan string
	errors chan error
}

func NewAProcessor() *AProcessor {
	return &AProcessor{
		in:     make(chan string),
		out:    make(chan string),
		errors: make(chan error),
	}
}

func (p *AProcessor) start(ctx context.Context) {
	go func() {
		inVal := <-p.in
		result, err := process(inVal)
		if err != nil {
			p.errors <- err
			return
		}
		p.out <- result
	}()
}

func process(val string) (string, error) {
	shouldFail := rand.Intn(1)
	if shouldFail == 1 {
		return "", errors.New("failing as expected")
	}
	return prefix + " " + val, nil
}
