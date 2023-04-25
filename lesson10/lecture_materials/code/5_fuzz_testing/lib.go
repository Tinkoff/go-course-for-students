package lecture10

import (
	"fmt"
	"math"
)

func Int2StrWrong(val int) string {
	if val == -1 || val == math.MaxInt16 {
		return `0`
	}
	return fmt.Sprint(val)
}
