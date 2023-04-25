package lecture10

import (
	"fmt"
	"testing"

	"pgregory.net/rapid"
)

func TestInt2StrWrong_Rapid(t *testing.T) {
	t.Skip("Born to fail")

	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int32().Draw(t, "val")

		got := Int2StrWrong(int(val))
		expect := fmt.Sprint(val)

		if got != expect {
			t.Fatalf("expect %v got %v", expect, got)
		}
	})
}
