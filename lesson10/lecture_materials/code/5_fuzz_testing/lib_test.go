package lecture10

import (
	"fmt"
	"testing"
)

func FuzzInt2StrWrong_Fuzz(f *testing.F) {
	testcases := []int{90, 1000}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, s int) {
		got := Int2StrWrong(s)
		expect := fmt.Sprintf("%d", s)

		if got != expect {
			t.Errorf("For (%d) Expect: %s, but got: %s", s, expect, got)
		}
	})
}
