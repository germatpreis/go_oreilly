package format

import (
	"ch10/internal"
	"fmt"
)

func Number(num int) string {
	internal.Doubler(2)
	return fmt.Sprintf("The number is %d", num)
}
