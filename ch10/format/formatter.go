package format

import "fmt"
import "ch10/internal"

func Number(num int) string {
	internal.Doubler(2)
	return fmt.Sprintf("The number is %d", num)
}
