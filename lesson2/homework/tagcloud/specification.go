package tagcloud

import "math"

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	tagLink map[string]int
	tagList []TagStat
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
	d.tagLink = make(map[string]int)
	return *d
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
// TODO: You decide whether receiver should be a pointer or a value
func (d *TagCloud) AddTag(tag string) {
	// TODO: Implement this
	id := d.tagLink[tag]
	if len(d.tagList) == 0 || d.tagList[id].Tag != tag {
		d.tagList = append(d.tagList, TagStat{tag, 1})
		d.tagLink[tag] = len(d.tagList) - 1
	} else {
		d.tagList[id].OccurrenceCount += 1
		if id > 0 && d.tagList[id-1].OccurrenceCount < d.tagList[id].OccurrenceCount {
			d.tagLink[tag], d.tagLink[d.tagList[id-1].Tag] = d.tagLink[d.tagList[id-1].Tag], d.tagLink[tag]
			d.tagList[id-1], d.tagList[id] = d.tagList[id], d.tagList[id-1]
		}
	}

}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
// TODO: You decide whether receiver should be a pointer or a value
func (d *TagCloud) TopN(n int) []TagStat {
	// TODO: Implement this
	return d.tagList[:int(math.Min(float64(n), float64(len(d.tagList))))]
}
