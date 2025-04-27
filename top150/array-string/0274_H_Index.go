package arraystring

import "sort"

func hIndex(citations []int) int {

	sort.Ints(citations)

	var hIndex int

	size := len(citations)
	for i := 0; i < size; i++ {
		if citations[i] > 0 {
			currhIndex := min(citations[i], size-i)
			if currhIndex > hIndex {
				hIndex = currhIndex
			}
		}
	}
	return hIndex
}

// 0 1 3 5 6

// if citations[i] > 0{
//     currHIndex := 1
//     atLeastCitations = size-i-1

//     hIndex := min(citations[i],currHIndex+atLeastCitations)
// }
