package setup_up_tear_down

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("(1) Use stuff set up in TestMain ", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("(2): Use stuff set up in TestMain ", testTime)
}
