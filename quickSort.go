package main

import "log"

func QuickSort(arr []int, left int, right int) {
	middle := left + (right-left)/2
	pivot := arr[middle]
	i := left
	j := right

	if left >= right {
		log.Fatal()
	}

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--

		}
	}
	if left < j {
		QuickSort(arr, left, j)
	}
	if right > i  {
		QuickSort(arr, i, right)
	}
}
