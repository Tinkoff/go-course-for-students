package tagcloud

import "sort"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tagTable map[string]int
	// TODO: add fields if necessary
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
// TODO: You decide whether this function should return a pointer or a value
func New() TagCloud {
	// TODO: Implement this
	d := new(TagCloud)
	d.tagTable = make(map[string]int)
	return *d
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
// TODO: You decide whether receiver should be a pointer or a value
func (d *TagCloud) AddTag(tag string) {
	// TODO: Implement this
	d.tagTable[tag] += 1
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
// TODO: You decide whether receiver should be a pointer or a value
func (d *TagCloud) TopN(n int) []TagStat {
	// TODO: Implement this
	topList := make([]TagStat, 0, len((*d).tagTable))

	for key, val := range (*d).tagTable {
		topList = append(topList, TagStat{key, val})
	}
	sort.Slice(topList, func(i, j int) bool {
		return topList[i].OccurrenceCount > topList[j].OccurrenceCount
	})
	for len(topList) > n {
		topList = topList[:len(topList)-1]
	}
	return topList
}
