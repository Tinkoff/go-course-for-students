package tagcloud

import "math"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	SortedTagStats []TagStat
	TagId          map[string]int
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

func New() *TagCloud {
	return &TagCloud{
		SortedTagStats: []TagStat{},
		TagId:          map[string]int{},
	}
}

func (tc *TagCloud) AddTag(tag string) {
	if id, hasValue := tc.TagId[tag]; hasValue {
		tc.SortedTagStats[id].OccurrenceCount++
		cnt := tc.SortedTagStats[id].OccurrenceCount
		isSwapped := false

		for i := id - 1; i >= 0; i-- {
			if tc.SortedTagStats[i].OccurrenceCount >= cnt {
				newId := i + 1
				tc.TagId[tag] = newId
				tc.TagId[tc.SortedTagStats[newId].Tag] = id
				a, b := tc.SortedTagStats[newId], tc.SortedTagStats[id]
				b, a = a, b
				tc.SortedTagStats[newId] = a
				tc.SortedTagStats[id] = b
				isSwapped = true
				break
			}
		}
		if !isSwapped {
			a, b := tc.SortedTagStats[0], tc.SortedTagStats[id]
			b, a = a, b
			tc.SortedTagStats[0] = a
			tc.SortedTagStats[id] = b
		}
	} else {
		tc.TagId[tag] = len(tc.SortedTagStats)
		tc.SortedTagStats = append(tc.SortedTagStats, TagStat{Tag: tag, OccurrenceCount: 1})
	}
}

func (tc *TagCloud) TopN(n int) []TagStat {
	n = int(math.Min(float64(n), float64(len(tc.SortedTagStats))))
	return tc.SortedTagStats[:n]
}
