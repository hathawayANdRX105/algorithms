package array_and_string

func Merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size < 2 {
		return intervals
	}

	quickSortForMerge(intervals, 0, size-1)

	var uniqueIndex int
	uniqueIntervals := make([][]int, size)
	uniqueIntervals[0] = intervals[0]
	
	for i := 1; i < size; i++ {

		if intervals[i-1][1] >= intervals[i][0] {
			// range merge
			// choose the smaller start range num
			intervals[i][0] = intervals[i-1][0]

			// choose the bigger end range num
			if intervals[i-1][1] > intervals[i][1]{
				intervals[i][1] = intervals[i-1][1]
			}
			
			// update merge range array
			uniqueIntervals[uniqueIndex] = intervals[i]

			// look for the next range whether need to be merge
		} else {
			// stop merge for uniqueIndex and shift next one to keep looking.
			uniqueIndex++
			uniqueIntervals[uniqueIndex] = intervals[i]
		}
	}

	return uniqueIntervals[:uniqueIndex+1]
}

func quickSortForMerge(intervals [][]int, lo, hi int) {
	if hi <= lo {
		return
	}

	pivot := partitionForMerge(intervals, lo, hi)
	quickSortForMerge(intervals, lo, pivot-1)
	quickSortForMerge(intervals, pivot+1, hi)
}

// partition according the start value in range
func partitionForMerge(arr [][]int, lo, hi int) int {
	cmpValue := arr[lo][0]

	i, j := lo, hi+1

	for {
		for i++; arr[i][0] < cmpValue; i++ {
			if i == hi {
				break
			}
		}

		for j--; arr[j][0] > cmpValue; j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		//exchange value of index i, j
		arr[i], arr[j] = arr[j], arr[i]

	}

	//exchange cmpValue with value of index j
	arr[lo], arr[j] = arr[j], arr[lo]

	return j
}
