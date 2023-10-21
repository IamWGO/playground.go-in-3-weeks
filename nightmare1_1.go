// N 1.1 Median of Two Sorted Arrays
// Given two sorted arrays nums1 and nums2 of size m and n respectively,
// return the median of the two sorted arrays.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getInputNumber(text string) int {
	for {
		fmt.Print(text)
		input := ""
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Empty!! input some number")
			continue
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("The input is not a valid integer.")
			continue
		}
		return number
	}
}

func inputNumbers(numbers []int, loop int) []int{
	for i:=0; i<loop; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	return numbers
}

func getMedianOfSortedArrays(combined []int) float64 {
	middleIndex := (len(combined) / 2)
	middleItem :=  combined[middleIndex] 

	medianOfSortedArrays := 0.0

	if (len(combined) % 2) == 0 {
		sum := combined[middleIndex -1] + combined[middleIndex]
		medianOfSortedArrays = float64(sum)/2
	} else {
		medianOfSortedArrays = float64(middleItem)
	}

	return medianOfSortedArrays
}


func main() {

	maxSet1 := getInputNumber("\nAmount of number list set 1 : ")
	numbers1 := []int{}
	numbers1 = inputNumbers(numbers1, maxSet1)

	maxSet2 := getInputNumber("\nAmount of number list set 2 : ")
	numbers2 := []int{}
	numbers2 = inputNumbers(numbers2, maxSet2)
	
	combined := append(numbers1, numbers2...)
	sort.Ints(combined)

	medianOfSortedArrays := getMedianOfSortedArrays(combined)

	fmt.Println("\nCombine arrays")
	fmt.Println(combined)
	fmt.Println("\nMedian of Two Sorted Arrays")
	fmt.Println(strconv.FormatFloat(medianOfSortedArrays, 'f', -1, 64))
	
	fmt.Printf("-- End program Nightmare1.1 --")
}


// The overall run time complexity should be O(log (m+n)).

// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.

// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10**6 <= nums1[i], nums2[i] <= 10**6