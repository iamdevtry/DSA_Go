package main

// Given a list of intervals, merge all overlapping intervals.
// Input: {[1, 4], [3, 6], [8, 10]}, Output: {[1, 6], [8, 10]}

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type IntervalList []Interval

func (intervals IntervalList) Len() int {
	return len(intervals)
}

func (intervals IntervalList) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func (intervals IntervalList) Less(i, j int) bool {
	return intervals[i].Start < intervals[j].Start
}

func mergeIntervals(intervals []Interval) []Interval {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	// sort intervals by start time
	sort.Sort(IntervalList(intervals))

	merged := make([]Interval, 0, n)
	merged = append(merged, intervals[0])

	for i := 1; i < n; i++ {
		if intervals[i].Start <= merged[len(merged)-1].End {
			if intervals[i].End > merged[len(merged)-1].End {
				merged[len(merged)-1].End = intervals[i].End
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

func main() {
	intervals := []Interval{
		{Start: 1, End: 4},
		{Start: 3, End: 6},
		{Start: 8, End: 10},
	}
	merged := mergeIntervals(intervals)
	fmt.Println(merged)
}
