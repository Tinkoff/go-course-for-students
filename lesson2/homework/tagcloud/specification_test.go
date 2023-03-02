package tagcloud_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"lecture02_homework/tagcloud"
	"testing"
)

func TestEmptyTagCloud(t *testing.T) {
	tc := tagcloud.New()
	topN := tc.TopN(1000)
	assert.Len(t, topN, 0, "empty tag cloud returned non-empty topN elements %v", topN)
}

func TestTopNGreaterThanCloudSize(t *testing.T) {
	tc := tagcloud.New()
	tc.AddTag("t1")

	requestCount := 10
	topN := tc.TopN(requestCount)
	assert.Len(t, topN, 1, "TopN(%d) returned array with invalid size: %v", requestCount, topN)
}

func TestHappyPath(t *testing.T) {
	tc := tagcloud.New()

	tc.AddTag("single-occurrence")
	tc.AddTag("multiple-occurrence")
	tc.AddTag("multiple-occurrence")

	top := tc.TopN(1)
	assert.Len(t, top, 1, "TopN(1) returned %d elements", len(top))

	if assert.Equal(t, "multiple-occurrence", top[0].Tag) {
		assert.Equal(t, 2, top[0].OccurrenceCount)
	}
}

func TestTopN(t *testing.T) {
	tc := tagcloud.New()
	size := 1000
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			tc.AddTag(fmt.Sprintf("%d", i))
		}
	}

	validateTopN := func(n int) {
		topN := tc.TopN(n)
		assert.Len(t, topN, n)

		for i, el := range topN {
			value := size - i - 1
			tagName := fmt.Sprintf("%d", value)
			assert.Equal(t, tagName, el.Tag, "TopN(%d) returned elements in wrong order (bad tag name at %d): %v", n, i, topN)
			assert.Equal(t, value, el.OccurrenceCount, "TopN(%d) returned elements in wrong order (bad occurrence count at %d): %v", n, i, topN)
		}
	}

	for i := 0; i < size; i++ {
		validateTopN(i)
	}
}

func TestTopNWithRepeatedOccurrence(t *testing.T) {
	tc := tagcloud.New()
	tc.AddTag("t1")
	tc.AddTag("t2")
	tc.AddTag("t3")
	tc.AddTag("t4")

	requestCount := 3
	topN := tc.TopN(requestCount)
	assert.Len(t, topN, requestCount)

	distinctMap := make(map[string]struct{})
	for _, v := range topN {
		assert.Equal(t, 1, v.OccurrenceCount)
		distinctMap[v.Tag] = struct{}{}
	}

	assert.Len(t, distinctMap, requestCount, "TopN(%d) returned array with non-distinct tags: %v", requestCount, topN)
}
