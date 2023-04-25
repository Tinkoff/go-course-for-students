package lecture10

import "testing"

var BenchSink int // make sure compiler cannot optimize away benchmarks

func BenchmarkInt2Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Int2Str(i)
		BenchSink += len(res)
	}
}

func BenchmarkInt2StrFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Int2StrFast(i)
		BenchSink += len(res)
	}
}

func BenchmarkInt2ByteSlice(b *testing.B) {
	var res []byte
	for i := 0; i < b.N; i++ {
		res = Int2ByteSlice(i, res[:0])
		BenchSink += len(res)
	}
}
