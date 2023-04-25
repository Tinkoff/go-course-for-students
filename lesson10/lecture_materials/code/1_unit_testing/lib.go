package lecture10

import (
	"fmt"
	"math"
)

func Int2Str(val int) string {
	return fmt.Sprint(val)
}

func Int2StrWrong(val int) string {
	if val == -1 || val == math.MaxInt16 {
		return `0`
	}
	return fmt.Sprint(val)
}

func Str2Int(val string) (res int) {
	_, _ = fmt.Sscan(val, &res)
	return
}
