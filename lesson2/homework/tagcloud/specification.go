package tagcloud

import "sort"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tagsStat map[string]int32
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
// TODO: You decide whether this function should return a pointer or a value
func New() TagCloud {
	m := make(map[string]int32)
	return TagCloud{tagsStat: m}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
// TODO: You decide whether receiver should be a pointer or a value
func (s *TagCloud) AddTag(tag string) {
	s.tagsStat[tag]++
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
// TODO: You decide whether receiver should be a pointer or a value
func (s *TagCloud) TopN(n int) []TagStat {
	tagsSlice := make([]TagStat, 0, len(s.tagsStat))
	for str, cnt := range s.tagsStat {
		tagsSlice = append(tagsSlice, TagStat{str, int(cnt)})
	}

	sort.Slice(tagsSlice, func(i, j int) bool {
		if tagsSlice[i].OccurrenceCount == tagsSlice[j].OccurrenceCount {
			return tagsSlice[i].Tag < tagsSlice[j].Tag
		}
		return tagsSlice[i].OccurrenceCount > tagsSlice[j].OccurrenceCount
	})

	if n > len(s.tagsStat) {
		n = len(s.tagsStat)
	}

	result := make([]TagStat, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, tagsSlice[i])
	}
	return result
}
