package cleanup

import (
	"errors"
	"os"
	"testing"
)

func createFile(t *testing.T) (_ string, err error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}
