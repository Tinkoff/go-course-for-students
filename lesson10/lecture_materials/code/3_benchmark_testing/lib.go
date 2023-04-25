package lecture10

import (
	"fmt"
	"strconv"
)

func Int2Str(val int) string {
	return fmt.Sprint(val)
}

func Int2StrFast(val int) string {
	return strconv.Itoa(val)
}

func Int2ByteSlice(val int, dst []byte) []byte {
	return strconv.AppendInt(dst, int64(val), 10)
}
