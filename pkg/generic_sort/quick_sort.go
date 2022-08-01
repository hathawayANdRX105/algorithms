package generic_sort

import (
	"math/rand"
	"time"
)



// partition ...
func partition(arr []int, lo, hi int) int {
	cmpValue := arr[lo]

	i, j := lo, hi+1

	for {
		for i++; arr[i] < cmpValue; i++ {
			if i == hi {
				break
			}
		}

		for j--; arr[j] > cmpValue; j-- {
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

// quickSort hide logic for quick sort.
func quickSort(arr []int, lo, hi int) {
	if hi <= lo {
		return
	}

	pivot := partition(arr, lo, hi)
	quickSort(arr, lo, pivot-1)
	quickSort(arr, pivot+1, hi)
}

// QuickSort use partition to split arrary and quick sort
func QuickSortByPartition(arr []int) {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	quickSort(arr, 0, len(arr)-1)
	
}

// quick3Way ...
func quick3Way(arr []int, lo, hi int) {
	if hi <= lo {
		return
	}

	lt, i, gt := lo, lo+1, hi

	cmpValue := arr[lo]

	for i <= gt {
		cmp := arr[i] - cmpValue

		if cmp < 0 {
			// exchage the value of index i to point of index lt
			// because the point of index lt point to the same as cmpValue
			// so i increme one.
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
			lt++
		} else if cmp > 0 {
			// exchange the value greater than cmpValue to point of index gt
			// but origin value of index gt hasn't compared.
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			// the value is same as cmpValue, just pass it.
			i++
		}
	}

	// the range between lt and gt include the value same as cmpValue.
	quick3Way(arr, lo, lt-1)
	quick3Way(arr, gt+1, hi)
}

func QuickSortBy3Way(arr []int) {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	quick3Way(arr, 0, len(arr)-1)
}
