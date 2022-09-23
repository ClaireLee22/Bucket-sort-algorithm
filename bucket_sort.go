package main

import "fmt"

func bucketSort(array []int) []int {
	maxValue := max(array)
	minValue := min(array)
	n := len(array)

	numOfBucket := (maxValue-minValue)/n + 1

	// create empty buckets
	buckets := [][]int{}
	for i := 0; i < numOfBucket; i++ {
		buckets = append(buckets, []int{})
	}

	// allocate elements to their corresponding bucket
	for _, element := range array {
		bucketIdx := (element - minValue) / n
		buckets[bucketIdx] = append(buckets[bucketIdx], element)
	}

	fmt.Println("unsorted buckets", buckets)

	// sort individual buckets using insertion sort
	for _, bucket := range buckets {
		insertionSort(bucket)
	}

	fmt.Println("sorted buckets", buckets)

	// combine sorted buckets to a sorted array
	sortedIdx := 0
	for _, bucket := range buckets {
		for idx := 0; idx < len(bucket); idx++ {
			array[sortedIdx] = bucket[idx]
			sortedIdx += 1
		}
	}

	return array
}

func max(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	return maxNum
}

func min(nums []int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}

	return minNum
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			swap(j-1, j, array)
			j -= 1
		}
	}
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	arrays := [][]int{
		{15, 5, 2, 9, 8, 4, 22},
		{15, 8, 4, -9, 5, -2, 22},
	}

	for _, array := range arrays {
		fmt.Println("unsorted array", array)
		fmt.Println("sorted array", bucketSort(array))
		fmt.Println()
	}
}

/* output:
unsorted array [15 5 2 9 8 4 22]
unsorted buckets [[5 2 8 4] [15 9] [22]]
sorted buckets [[2 4 5 8] [9 15] [22]]
sorted array [2 4 5 8 9 15 22]

unsorted array [15 8 4 -9 5 -2 22]
unsorted buckets [[-9] [4 -2] [8 5] [15] [22]]
sorted buckets [[-9] [-2 4] [5 8] [15] [22]]
sorted array [-9 -2 4 5 8 15 22]
*/
