package executor

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	sleepPerStage = time.Millisecond * 100
	fault         = sleepPerStage / 2
)

func TestPipeline(t *testing.T) {
	// Stage generator
	g := func(_ string, f func(v any) any) Stage {
		return func(in In) Out {
			out := make(chan any)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage)
					out <- f(v)
				}
			}()
			return out
		}
	}

	t.Run("one stage one value", func(t *testing.T) {
		in := make(chan any)
		data := []int{1}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]any, 0, len(data))
		start := time.Now()
		for s := range ExecutePipeline(context.Background(), in, g("Multiplier (* 100)", func(v any) any { return v.(int) * 100 })) {
			result = append(result, s)
		}

		elapsed := time.Since(start)

		require.Equal(t, []any{100}, result)
		require.Less(t, int64(elapsed), // ~0.1s for processing 1 value in 1 stages
			int64(sleepPerStage)+int64(fault))
	})

	stages := []Stage{
		g("Dummy", func(v any) any { return v }),
		g("Multiplier (* 2)", func(v any) any { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v any) any { return v.(int) + 100 }),
		g("Stringifier", func(v any) any { return strconv.Itoa(v.(int)) }),
	}

	t.Run("multiple stages one value", func(t *testing.T) {
		in := make(chan any)
		data := []int{1}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, len(data))
		start := time.Now()
		for s := range ExecutePipeline(context.Background(), in, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Equal(t, []string{"102"}, result)
		require.Less(t, int64(elapsed), // ~0.4s for processing 1 valuee in 4 stages (100ms every) concurrently
			int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))
	})

	t.Run("multiple stages multiple values", func(t *testing.T) {
		in := make(chan any)
		data := []int{1, 2, 3, 4, 5}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, len(data))
		start := time.Now()
		for s := range ExecutePipeline(context.Background(), in, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Equal(t, []string{"102", "104", "106", "108", "110"}, result)
		require.Less(t, int64(elapsed), // ~0.8s for processing 5 values in 4 stages (100ms every) concurrently
			int64(sleepPerStage)*int64(len(stages)+len(data)-1)+int64(fault))
	})

	t.Run("ctx cancel case", func(t *testing.T) {
		// Abort after 200ms
		abortDur := sleepPerStage * 2

		in := make(chan any)
		ctx, cancel := context.WithTimeout(context.Background(), abortDur)

		defer cancel()

		stopGenCh := make(chan struct{})
		defer close(stopGenCh)

		// gen values forever
		go func() {
			defer close(in)

			i := 0
			for {
				select {
				case <-stopGenCh:
					return
				default:
					in <- i
					i++
				}
			}
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(ctx, in, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Len(t, result, 0)
		require.Less(t, int64(elapsed), int64(abortDur)+int64(fault))
	})
}
